module Style exposing (..)

import Css


h_x_px : Int -> Css.Style
h_x_px height =
    Css.property "height" <| String.fromInt height ++ "px"
