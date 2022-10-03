module Error exposing (..)

import Html.Styled exposing (Html, div, text)
import Html.Styled.Attributes as Attr exposing (..)
import Play exposing (Msg)
import Tailwind.Utilities exposing (..)


view : String -> Html Msg
view error =
    div
        [ Attr.css [ flex_none, font_sans, text_lg, text_red_500 ]
        , Attr.id "result-error"
        ]
        [ text error ]
