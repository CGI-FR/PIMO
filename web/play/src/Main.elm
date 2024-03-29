module Main exposing (init, main, update, view)

import Browser
import Css.Global
import Error
import Examples
import Header exposing (view)
import Html.Styled as Html exposing (..)
import Html.Styled.Attributes as Attr exposing (..)
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
      }
    , Cmd.none
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

        GotMaskedData output ->
            ( { model
                | output = output
                , status = Success
                , error = ""
              }
            , updateOutputEditor output
            )

        GotFlowData flow ->
            let
                cmd =
                    case model.maskingView of
                        GraphView ->
                            updateFlow flow

                        _ ->
                            Cmd.none
            in
            ( { model
                | flow = flow
                , status = Success
                , error = ""
              }
            , cmd
            )

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
                cmd =
                    case maskingView of
                        GraphView ->
                            updateFlow model.flow

                        _ ->
                            Cmd.none
            in
            ( { model | maskingView = maskingView }, cmd )

        Error errorMessage ->
            (
                { model
                | error = errorMessage
                ,  status = Success
                }
            , Cmd.none )



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
        , outputUpdater mapOutputUpdater
        , flowUpdater mapFlowUpdater
        , errorUpdater mapErrorUpdater
        ]


mapFlowUpdater : JD.Value -> Msg
mapFlowUpdater flow =
    case JD.decodeValue JD.string flow of
        Ok data ->
            GotFlowData data

        Err errorMessage ->
            Error (JD.errorToString errorMessage)


mapOutputUpdater : JD.Value -> Msg
mapOutputUpdater outputJson =
    case JD.decodeValue JD.string outputJson of
        Ok output ->
            GotMaskedData output

        Err errorMessage ->
            Error (JD.errorToString errorMessage)



mapErrorUpdater : JD.Value -> Msg
mapErrorUpdater errorJson =
    case JD.decodeValue JD.string errorJson of
        Ok error ->
            Error error

        Err errorMessage ->
            Error (JD.errorToString errorMessage)

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
        [ pimoMask sandbox
        ]
