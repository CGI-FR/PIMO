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
        [ Attr.css [ flex, flex_col ]
        ]
        [button
            [ Attr.id "dropdownDefault"
            , Attr.attribute "data-dropdown-toggle" "dropdown"
            , css
                [ Tw.text_white
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
                ]
            , Attr.type_ "button"
            , onCustomClick <| ChangeDropdownView Open
            ]
            [ text "Dropdown button", svg
                [ SvgAttr.css
                    [ Tw.ml_2
                    , Tw.w_4
                    , Tw.h_4
                    ]
                , Attr.attribute "aria-hidden" "true"
                , SvgAttr.fill "none"
                , SvgAttr.stroke "currentColor"
                , SvgAttr.viewBox "0 0 24 24"
                ]
                [ Svg.path
                    [ SvgAttr.strokeLinecap "round"
                    , SvgAttr.strokeLinejoin "round"
                    , SvgAttr.strokeWidth "2"
                    , SvgAttr.d "M19 9l-7 7-7-7"
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
                    , Tw.text_gray_700
                    ]
                , Attr.attribute "aria-labelledby" "dropdownDefault"
                ]
                [ li []
                    [ a
                        [ Attr.href "#"
                        , css
                            [ Tw.block
                            , Tw.py_2
                            , Tw.px_4
                            , Css.hover
                                [ Tw.bg_gray_100
                                ]
                            ]
                        ]
                        [ text "Share : Copy link" ]
                    ]
                , li []
                    [ a
                        [ Attr.href "#"
                        , css
                            [ Tw.block
                            , Tw.py_2
                            , Tw.px_4
                            , Css.hover
                                [ Tw.bg_gray_100
                                ]
                            ]
                        ]
                        [ text "Export as Venom Test" ]
                    ]
                ]
            ]
        ]
