module Masking exposing (..)

import Css
import Html.Styled exposing (..)
import Html.Styled.Attributes as Attr exposing (..)
import Play exposing (..)
import Style exposing (h_x_px, onCustomClick)
import Tailwind.Breakpoints as Bp
import Tailwind.Utilities exposing (..)


selected : Attribute msg
selected =
    css
        [ inline_block
        , p_1
        , text_blue_600
        , bg_gray_100
        , rounded_t_lg
        ]


notSelected : Attribute msg
notSelected =
    css
        [ inline_block
        , p_1
        , rounded_t_lg
        , Css.hover
            [ text_gray_600
            , bg_gray_50
            ]
        ]


tabView : MaskingView -> List (Html Msg)
tabView maskingView =
    [ ul
        [ css
            [ flex
            , flex_wrap
            , text_sm
            , font_medium
            , text_center
            , text_gray_500
            , border_b
            , border_gray_200
            ]
        ]
        [ li
            [ css
                [ mr_2
                ]
            ]
            [ a
                [ Attr.href "#"
                , case maskingView of
                    YamlView ->
                        selected

                    _ ->
                        notSelected
                , onCustomClick <| ChangeMaskingView YamlView
                ]
                [ text "YAML" ]
            ]
        , li
            [ css
                [ mr_2
                ]
            ]
            [ a
                [ Attr.href "#"
                , case maskingView of
                    GraphView ->
                        selected

                    _ ->
                        notSelected
                , onCustomClick <| ChangeMaskingView GraphView
                ]
                [ text "Graph" ]
            ]
        ]
    ]


view : MaskingView -> Html Msg
view maskingView =
    div
        [ Attr.css [ flex, flex_col ]
        ]
        (div
            [ Attr.css [ flex_none, font_sans, text_xl, pb_2 ]
            ]
            [ text "Masking Configuration" ]
            :: tabView maskingView
            ++ [ div
                    [ Attr.css [ grow, shadow_lg, h_x_px 300, Bp.md [ h_x_px 600 ] ]
                    , Attr.id "editor-yaml"
                    , Attr.hidden (maskingView /= YamlView)
                    ]
                    []
               , div
                    [ Attr.css [ grow, shadow_lg, h_x_px 300, Bp.md [ h_x_px 600 ] ]
                    , Attr.id "flowchart"
                    ,  Attr.hidden (maskingView /= GraphView)
                    ]
                    []
               ]
        )
