module Header exposing (..)

import Html.Styled exposing (..)
import Html.Styled.Attributes as Attr exposing (..)
import Play exposing (Msg(..), init_sandbox)
import Style exposing (onCustomClick)
import Tailwind.Utilities as Tw exposing (..)


view : String -> Html Msg
view version =
    div
        [ Attr.css [ mb_5, flex, space_x_4, items_baseline, bg_black, text_white ] ]
        [ div
            [ Attr.css [ m_5, font_sans, text_4xl, font_bold ]
            ]
            [ a
                [ Attr.id "reset-link"
                , Attr.css [ cursor_pointer ]
                , onCustomClick <| UpdateMaskingAndInput init_sandbox
                ]
                [ text "PIMO Play" ]
            ]
        , div
            [ Attr.css [ text_slate_200, text_sm ]
            ]
            [ text "A playground for "
            , a
                [ Attr.href ("https://github.com/CGI-FR/PIMO/tree/" ++ version)
                , Attr.target "_blank"
                , Attr.rel "noopener noreferrer"
                , Attr.css [ Tw.no_underline ]
                ]
                [ text ("pimo " ++ version) ]
            ]
        , div
            [ Attr.css [ grow ]
            ]
            []
        ]
