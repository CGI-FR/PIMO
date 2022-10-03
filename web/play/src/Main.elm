module Main exposing (init, main, update, view)

import Browser
import Css
import Css.Global
import Debug exposing (toString)
import Error
import Examples
import Header exposing (view)
import Html.Styled as Html exposing (..)
import Html.Styled.Attributes as Attr exposing (..)
import Http
import Http.Detailed
import Json.Decode as JD
import Json.Encode as JE
import OutputPanel
import Play exposing (..)
import Ports exposing (..)
import Style exposing (h_x_px)
import Tailwind.Breakpoints as Breakpoints
import Tailwind.Utilities as Tw exposing (..)


init : () -> ( Model, Cmd Msg )
init () =
    ( { version = "master"
      , sandbox = { masking = "masking", input = "{}" }
      , output = "{}"
      , error = ""
      , status = Loading
      }
    , Cmd.batch
        [ initMaskingEditor "masking"
        , initInputEditor "{}"
        , initOutputEditor "{}"
        ]
    )



-- ---------------------------
-- UPDATE
-- ---------------------------


update : Msg -> Model -> ( Model, Cmd Msg )
update message model =
    case message of
        UpdateInput input ->
            let
                newModel =
                    Loading
                        |> asStatusIn
                            (input
                                |> asInputIn model.sandbox
                                |> asSandboxIn model
                            )
            in
            ( newModel, maskRequest newModel )

        UpdateMasking masking ->
            let
                newModel =
                    Loading
                        |> asStatusIn
                            (masking
                                |> asMaskingIn model.sandbox
                                |> asSandboxIn model
                            )
            in
            ( newModel
            , maskRequest newModel
            )

        Refresh ->
            let
                newModel =
                    Loading
                        |> asStatusIn model
            in
            ( newModel
            , maskRequest newModel
            )

        GotMaskedData result ->
            case result of
                Ok ( _, output ) ->
                    ( { model
                        | output = output
                        , status = Success
                        , error = ""
                      }
                    , updateOutputEditor output
                    )

                Err error ->
                    let
                        errorMessage =
                            case error of
                                Http.Detailed.BadStatus _ body ->
                                    body

                                _ ->
                                    "Server Error"
                    in
                    ( { model | error = errorMessage, status = Failure }, Cmd.none )

        UpdateMaskingAndInput sandbox ->
            let
                newModel =
                    Loading
                        |> (asStatusIn <| (sandbox |> asSandboxIn model))
            in
            ( newModel
            , Cmd.batch [ updateMaskingEditor sandbox.masking, updateInputEditor sandbox.input, maskRequest newModel ]
            )

        Error errorMessage ->
            ( { model | error = errorMessage }, Cmd.none )



-- ---------------------------
-- VIEW
-- ---------------------------
-- view : Model -> VirtualDom.Node Msg


view model =
    Html.toUnstyled <|
        node "body"
            []
            [ Css.Global.global Tw.globalStyles
            , Header.view model.version
            , div
                [ Attr.css [ Tw.px_4, Breakpoints.lg [ Tw.px_16 ], Breakpoints.md [ Tw.px_8 ] ] ]
                [ div
                    [ Attr.css [ Tw.grid, Breakpoints.sm [ Tw.grid_cols_2 ], Tw.gap_4, Breakpoints.md [ gap_8 ] ] --  "grid grid-cols-1 sm:grid-cols-2 gap-4 md:gap-8"
                    ]
                    [ div
                        [ Attr.css [ Tw.flex, Tw.flex_col ]
                        ]
                        [ div
                            [ Attr.css [ flex_none, font_sans, text_xl, pb_2 ]
                            ]
                            [ text "Masking Configuration" ]
                        , div
                            [ Attr.css [ grow, shadow_lg, h_x_px 300, Breakpoints.md [ h_x_px 600 ] ]
                            , Attr.id "editor-yaml"
                            ]
                            []
                        ]
                    , div
                        [ Attr.css [ flex, flex_col ]
                        ]
                        ([ div
                            [ Attr.css [ flex_none, font_sans, text_xl, pb_2 ]
                            ]
                            [ text "Input" ]
                         , div
                            [ Attr.css [ grow, shadow_lg, h_x_px 300, Breakpoints.md [ h_full ] ]
                            , Attr.id "editor-json"
                            ]
                            []
                         ]
                            ++ OutputPanel.view model.status
                        )
                    ]
                , Error.view model.error
                , Examples.view
                ]
            ]



-- ---------------------------
-- MAIN
-- ---------------------------


main : Program () Model Msg
main =
    Browser.document
        { init = init
        , update = update
        , view =
            \m ->
                { title = "PIMO Play !"
                , body = [ view m ]
                }
        , subscriptions = subscriptions
        }


subscriptions : Model -> Sub Msg
subscriptions _ =
    Sub.batch
        [ maskingUpdater UpdateMasking
        , inputUpdater UpdateInput
        , maskingAndinputUpdater mapMaskingAndinputUpdater
        ]


mapMaskingAndinputUpdater : JD.Value -> Msg
mapMaskingAndinputUpdater sandboxJson =
    case JD.decodeValue sandboxDecoder sandboxJson of
        Ok sandbox ->
            UpdateMaskingAndInput sandbox

        Err errorMessage ->
            Error (JD.errorToString errorMessage)


sandboxDecoder : JD.Decoder Sandbox
sandboxDecoder =
    JD.map2 Sandbox
        (JD.field "masking" JD.string)
        (JD.field "input" JD.string)


maskRequestEncoder : Model -> JE.Value
maskRequestEncoder model =
    JE.object
        [ ( "masking", JE.string model.sandbox.masking )
        , ( "data", JE.string model.sandbox.input )
        ]


maskRequest : Model -> Cmd Msg
maskRequest model =
    Http.post
        { url = "/play"
        , body = Http.jsonBody <| maskRequestEncoder model
        , expect = Http.Detailed.expectString GotMaskedData
        }
