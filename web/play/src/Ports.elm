port module Ports exposing (..)

import Json.Decode as JD



-- ---------------------------
-- PORTS
-- ---------------------------


port initMaskingEditor : String -> Cmd msg


port maskingUpdater : (String -> msg) -> Sub msg


port updateMaskingEditor : String -> Cmd msg


port initInputEditor : String -> Cmd msg


port inputUpdater : (String -> msg) -> Sub msg


port updateInputEditor : String -> Cmd msg


port initOutputEditor : String -> Cmd msg


port updateOutputEditor : String -> Cmd msg


port maskingAndinputUpdater : (JD.Value -> msg) -> Sub msg
