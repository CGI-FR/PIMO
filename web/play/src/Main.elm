module Main exposing (init, main, update, view)

import Browser
import Css.Global
import Error
import Examples
import Header exposing (view)
import Html.Styled as Html exposing (..)
import Html.Styled.Attributes as Attr exposing (..)
import Html.Styled.Events exposing (onClick)
import Http
import Http.Detailed
import Json.Decode as JD
import Json.Encode as JE
import Masking
import OutputPanel
import Play exposing (..)
import Ports exposing (..)
import Style exposing (h_x_px)
import Tailwind.Breakpoints as Breakpoints
import Tailwind.Utilities as Tw exposing (..)

init : String -> ( Model, Cmd Msg )
init version =
    ( { version = version
      , sandbox = init_sandbox
      , output = "{}"
      , error = ""
      , status = Loading
      , maskingView = YamlView
      , flow = ""
      , popupVisible = False
      }
    , fetchSecure
    )

secureDecoder : JD.Decoder Bool
secureDecoder =
    JD.field "secure" JD.bool

fetchSecure : Cmd Msg
fetchSecure =
    Http.get {url = "/options", expect = Http.expectJson ReceiveSecure secureDecoder}


-- ---------------------------
-- UPDATE
-- ---------------------------


update : Msg -> Model -> ( Model, Cmd Msg )
update message model =
    case message of
        UpdateInput input ->
            let
                newModel =
                    { model
                        | status = Loading
                        , sandbox = input |> asInputIn model.sandbox
                        , output = ""
                    }
            in
            ( newModel, maskRequest newModel.sandbox )

        UpdateMasking masking ->
            let
                newModel =
                    { model
                        | status = Loading
                        , sandbox = masking |> asMaskingIn model.sandbox
                        , output = ""
                    }
            in
            ( newModel
            , maskRequest newModel.sandbox
            )

        Refresh ->
            let
                newModel =
                    Loading
                        |> asStatusIn model
            in
            ( newModel
            , maskRequest newModel.sandbox
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


        GotFlowData result ->
            case result of
                Ok ( _, flow ) ->
                    let
                        cmd = case model.maskingView of
                            GraphView -> updateFlow flow
                            _ -> Cmd.none
                    in
                    ( { model
                        | flow = flow
                      }
                    , cmd
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
                    { model
                        | status = Loading
                        , sandbox = sandbox
                        , output = ""
                    }
            in
            ( newModel
            , Cmd.batch [ updateOutputEditor newModel.output, updateMaskingEditor sandbox.masking, updateInputEditor sandbox.input, maskRequest newModel.sandbox ]
            )

        ChangeMaskingView maskingView ->
            let
              cmd = case maskingView of
                GraphView -> updateFlow model.flow
                _ -> Cmd.none
            in
            ( { model | maskingView = maskingView },  cmd )

        Error errorMessage ->
            ( { model | error = errorMessage }, Cmd.none )

        ClosePopup ->
            ( { model | popupVisible = False }
            , Cmd.none
            )

        ReceiveSecure result ->
            case result of
                Ok secure ->
                    ( { model | popupVisible = secure }, Cmd.none )

                Err _ ->
                    ( model, Cmd.none )


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
            , popup model
            , div
                [ Attr.css [ Tw.px_4, Breakpoints.lg [ Tw.px_16 ], Breakpoints.md [ Tw.px_8 ] ] ]
                [ div
                    [ Attr.css [ Tw.grid, Breakpoints.sm [ Tw.grid_cols_2 ], Tw.gap_4, Breakpoints.md [ gap_8 ] ] --  "grid grid-cols-1 sm:grid-cols-2 gap-4 md:gap-8"
                    ]
                    [ Masking.view model.maskingView
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

popup : Model -> Html Msg
popup model =
    if model.popupVisible then
        div []
        [ div
            [ Attr.id "overlay"
            , css
                [ Tw.fixed
                , Tw.inset_0
                , Tw.bg_black
                , Tw.bg_opacity_50
                , Tw.z_40
                ]
            ]
            []
            , div
                [ css
                    [ Tw.fixed
                    , Tw.inset_0
                    , Tw.flex
                    , Tw.items_center
                    , Tw.justify_center
                    , Tw.z_50
                    ]
                ]
                [ div
                    [ css
                        [ Tw.bg_white
                        , Tw.p_6
                        , Tw.w_1over3
                        , Tw.rounded
                        , Tw.shadow_lg
                        ]
                    ]
                    [ h2
                        [ css
                            [ Tw.text_2xl
                            , Tw.mb_4
                            ]
                        ]
                        [ text "Welcome to PIMO Play !" ]
                    , p
                        [ css
                            [ Tw.text_gray_700
                            ]
                        ]
                        [ text "Please, never use real data on this service." ]
                    , button
                        [ onClick ClosePopup
                        , css
                            [ Tw.mt_4
                            , Tw.bg_blue_500
                            , Tw.text_white
                            , Tw.font_bold
                            , Tw.py_2
                            , Tw.px_4
                            , Tw.rounded
                            ]
                        ]
                        [ text "I agree" ]
                    ]
                ]
            ]
    else
        text ""



-- ---------------------------
-- MAIN
-- ---------------------------


main : Program String Model Msg
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


maskRequestEncoder : Sandbox -> JE.Value
maskRequestEncoder sandbox =
    JE.object
        [ ( "masking", JE.string sandbox.masking )
        , ( "data", JE.string sandbox.input )
        ]


flowRequestEncoder : String -> JE.Value
flowRequestEncoder masking =
    JE.object
        [ ( "masking", JE.string masking )
        ]


maskRequest : Sandbox -> Cmd Msg
maskRequest sandbox =
    Cmd.batch
        [ Http.post
            { url = "/play"
            , body = Http.jsonBody <| maskRequestEncoder sandbox
            , expect = Http.Detailed.expectString GotMaskedData
            }
        , Http.post
            { url = "/flow"
            , body = Http.jsonBody <| flowRequestEncoder sandbox.masking
            , expect = Http.Detailed.expectString GotFlowData
            }
        ]
