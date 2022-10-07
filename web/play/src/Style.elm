module Style exposing (..)

import Css
import Html.Styled exposing (..)
import Html.Styled.Events exposing (custom)
import Json.Decode as JD
import Play exposing (..)


onCustomClick : Msg -> Html.Styled.Attribute Msg
onCustomClick msg =
    custom
        "click"
        (JD.succeed
            { message = msg
            , stopPropagation = True
            , preventDefault = True
            }
        )


h_x_px : Int -> Css.Style
h_x_px height =
    Css.property "height" <| String.fromInt height ++ "px"
