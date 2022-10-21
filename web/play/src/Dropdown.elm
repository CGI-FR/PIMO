module Dropdown exposing (..)

import Css
import Play exposing (..)
import Html.Styled exposing (..)
import Html.Styled.Attributes as Attr exposing (..)
import Style exposing (onCustomClick)
import Svg.Styled as Svg exposing (svg, path)
import Svg.Styled.Attributes as SvgAttr
import Tailwind.Breakpoints as Bp
import Tailwind.Utilities as Tw exposing (..)

on : Attribute msg
on =
    css
        [ Tw.contents
        , Tw.z_10
        , Tw.w_44
        , Tw.bg_white
        , Tw.rounded
        , Tw.divide_y
        , Tw.divide_gray_100
        , Tw.shadow
        ]

off : Attribute msg
off =
    css
        [ Tw.hidden
        , Tw.z_10
        , Tw.w_44
        , Tw.bg_white
        , Tw.rounded
        , Tw.divide_y
        , Tw.divide_gray_100
        , Tw.shadow
        ]

view : DropdownView -> Html Msg
view dropdownView =
    div
        [ css
            [ Tw.flex
            , Tw.flex_row_reverse
            , Tw.mb_5
            , Tw.neg_mt_10
            , Tw.bg_black
            , Tw.text_white
            ]
        ]
        [button
            [ Attr.id "dropdownBottomButton"
            , Attr.attribute "data-dropdown-toggle" "dropdownBottom"
            , Attr.attribute "data-dropdown-placement" "bottom"
            , css
                [ Tw.mr_3
                , Tw.mb_3
                , Tw.text_white
                , Tw.bg_blue_700
                , Tw.font_medium
                , Tw.rounded_lg
                , Tw.text_sm
                , Tw.px_4
                , Tw.py_2_dot_5
                , Tw.text_center
                , Tw.inline_flex
                , Tw.items_center
                , Css.focus
                    [ Tw.ring_4
                    , Tw.outline_none
                    , Tw.ring_blue_300
                    ]
                , Css.hover
                    [ Tw.bg_blue_800
                    ]
                , Bp.md
                    [ Tw.mb_0
                    ]
                ]
            , Attr.type_ "button"
            , onCustomClick <| ChangeDropdownView Open
            ]
            [ text "Options", svg
                [ SvgAttr.css
                    [ Tw.ml_2
                    , Tw.w_4
                    , Tw.h_4
                    ]
                , Attr.attribute "aria-hidden" "true"
                , SvgAttr.fill "currentColor"
                , SvgAttr.viewBox "0 0 20 20"
                ]
                [ path
                    [ SvgAttr.fillRule "evenodd"
                    , SvgAttr.d "M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
                    , SvgAttr.clipRule "evenodd"
                    ]
                    []
                ]
            ]
        ,     {- Dropdown menu -}
            div
            [ Attr.id "dropdown"
            , case dropdownView of
                Open ->
                    on
                _ ->
                    off
            ]
            [ ul
                [ css
                    [ Tw.py_1
                    , Tw.text_sm
                    , Tw.text_white
                    ]
                , Attr.attribute "aria-labelledby" "dropdownBottomButton"
                ]
                [ li [Attr.id "share"]
                    [ a
                        [ Attr.href "#"
                        , css
                            [ Tw.block
                            , Tw.py_2
                            , Tw.px_4
                            , Css.hover
                                [ Tw.bg_blue_700
                                ]
                            ]
                        ]
                        [ text "Share : Copy link" ]
                    ]
                , li [Attr.id "venom"]
                    [ a
                        [ Attr.href "#"
                        , css
                            [ Tw.block
                            , Tw.py_2
                            , Tw.px_4
                            , Css.hover
                                [ Tw.bg_blue_700
                                ]
                            ]
                        ]
                        [ text "Export as Venom Test" ]
                    ]
                ]
            ]
        ]
