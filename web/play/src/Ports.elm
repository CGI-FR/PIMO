port module Ports exposing (..)

import Json.Decode as JD



-- ---------------------------
-- PORTS YAML Editor
-- ---------------------------


port maskingUpdater : (String -> msg) -> Sub msg


port updateMaskingEditor : String -> Cmd msg



-- ---------------------------
-- PORTS JSON input Editor
-- ---------------------------


port inputUpdater : (String -> msg) -> Sub msg


port updateInputEditor : String -> Cmd msg



-- ---------------------------
-- PORTS JSON output Editor
-- ---------------------------


port updateOutputEditor : String -> Cmd msg



-- ---------------------------
-- PORTS Loading example
-- ---------------------------


port maskingAndinputUpdater : (JD.Value -> msg) -> Sub msg



-- ---------------------------
-- PORTS updateFlow
-- ---------------------------


port updateFlow : String -> Cmd msg
