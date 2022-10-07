module Examples exposing (..)

import Css exposing (visited)
import Html.Styled exposing (..)
import Html.Styled.Attributes as Attr exposing (..)
import Tailwind.Breakpoints exposing (md)
import Tailwind.Utilities exposing (..)


view : Html msg
view =
    div
        [ Attr.css [ flex, flex_col, md [ flex_row ], mt_5, gap_2, justify_evenly ]
        ]
        [ div []
            [ h3
                [ Attr.css [ text_lg ]
                ]
                [ text "Data Generation" ]
            , a
                [ Attr.id "example-generation"
                , Attr.css [ text_blue_400, visited [ text_purple_400 ], text_sm, underline, block, cursor_pointer ]
                ]
                []
            ]
        , div []
            [ h3
                [ Attr.css [ text_lg ]
                ]
                [ text "Data Anonymization" ]
            , a
                [ Attr.id "example-anonymization"
                , Attr.css [ text_blue_400, visited [ text_purple_400 ], text_sm, underline, block, cursor_pointer ]
                ]
                []
            ]
        , div []
            [ h3
                [ Attr.css [ text_lg ]
                ]
                [ text "Data Pseudonymization" ]
            , a
                [ Attr.id "example-pseudonymization"
                , Attr.css [ text_blue_400, visited [ text_purple_400 ], text_sm, underline, block, cursor_pointer ]
                ]
                []
            ]
        , div []
            [ h3
                [ Attr.css [ text_lg ]
                ]
                [ text "Other" ]
            , a
                [ Attr.id "example-other"
                , Attr.css [ text_blue_400, visited [ text_purple_400 ], text_sm, underline, block, cursor_pointer ]
                ]
                []
            ]
        ]
