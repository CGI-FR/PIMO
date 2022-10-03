module OutputPanel exposing (..)

import Html exposing (Attribute)
import Html.Events exposing (custom)
import Html.Styled exposing (..)
import Html.Styled.Attributes as Attr exposing (class, href, id)
import Json.Decode as JD
import Play exposing (..)
import Style exposing (h_x_px)
import Svg.Styled as Svg exposing (path, svg)
import Svg.Styled.Attributes as SvgAttr
import Tailwind.Breakpoints as Breakpoints
import Tailwind.Utilities as Tw exposing (..)


onCustomClick : Msg -> Html.Styled.Attribute Msg
onCustomClick msg =
    Attr.fromUnstyled <|
        custom
            "click"
            (JD.succeed
                { message = msg
                , stopPropagation = True
                , preventDefault = True
                }
            )


spinner : Html msg
spinner =
    svg
        [ SvgAttr.id "refresh-spinner"
        , SvgAttr.css [ h_5, w_5, inline ]
        , SvgAttr.viewBox "0 0 14 14"
        ]
        [ Svg.g
            [ SvgAttr.fill "none"
            , SvgAttr.fillRule "evenodd"
            ]
            [ Svg.circle
                [ SvgAttr.cx "7"
                , SvgAttr.cy "7"
                , SvgAttr.r "6"
                , SvgAttr.stroke "#000"
                , SvgAttr.strokeOpacity ".1"
                , SvgAttr.strokeWidth "2"
                ]
                []
            , path
                [ SvgAttr.fill "#000"
                , SvgAttr.fillOpacity ".1"
                , SvgAttr.fillRule "nonzero"
                , SvgAttr.d "M7 0a7 7 0 0 1 7 7h-2a5 5 0 0 0-5-5V0z"
                ]
                []
            ]
        ]


refreshButton : Html Msg
refreshButton =
    a
        [ Attr.href "#"
        , Attr.id "refresh-button"
        , Attr.css [ text_2xl, font_medium ]
        , onCustomClick Refresh
        ]
        [ text "↻" ]


view : Status -> List (Html Msg)
view status =
    [ div
        [ Attr.css [ flex_none, font_sans, text_xl, py_2 ]
        , Attr.id "label-output"
        ]
        [ span [ Attr.css [ flex_none, font_sans, text_xl, py_2 ] ]
            [ text "Output " ]
        , case status of
            Loading ->
                spinner

            _ ->
                refreshButton
        ]
    , div
        [ Attr.css [ grow, shadow_lg, h_x_px 300, Breakpoints.md [ h_full ] ]
        , Attr.id "result-json"
        ]
        []
    ]
