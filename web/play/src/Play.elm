module Play exposing (..)

import Http
import Http.Detailed
import Json.Decode exposing (Error(..))



-- ---------------------------
-- MODEL
-- ---------------------------


init_sandbox : Sandbox
init_sandbox =
    { masking = """version: "1"
masking:
  - selector :
      jsonpath : "field_name"
    masks:
      - add: ""
     """
    , input = "{}"
    }


type alias Model =
    { version : String
    , sandbox : Sandbox
    , output : String
    , error : String
    , status : Status
    }


type alias Sandbox =
    { masking : String
    , input : String
    }


type Status
    = Success
    | Loading
    | Failure


type Msg
    = GotMaskedData (Result (Http.Detailed.Error String) ( Http.Metadata, String ))
    | UpdateMasking String
    | UpdateInput String
    | UpdateMaskingAndInput Sandbox
    | Refresh
    | Error String


asMaskingIn : Sandbox -> String -> Sandbox
asMaskingIn sandbox masking =
    { sandbox | masking = masking }


asInputIn : Sandbox -> String -> Sandbox
asInputIn sandbox input =
    { sandbox | input = input }


asSandboxIn : Model -> Sandbox -> Model
asSandboxIn model sandbox =
    { model | sandbox = sandbox }


asStatusIn : Model -> Status -> Model
asStatusIn model status =
    { model | status = status }
