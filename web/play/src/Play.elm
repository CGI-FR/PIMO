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
  - selector:
      jsonpath: "name"
    mask:
      randomChoiceInUri: "pimo://nameFR"
     """
    , input = """{
  "name": "Bill"
}"""
    }


type alias Model =
    { version : String
    , sandbox : Sandbox
    , output : String
    , error : String
    , status : Status
    , maskingView : MaskingView
    , flow :  String
    , popupVisible : Bool
    }


type alias Sandbox =
    { masking : String
    , input : String
    }


type Status
    = Success
    | Loading
    | Failure


type MaskingView
    = YamlView
    | GraphView


type Msg
    = GotMaskedData (Result (Http.Detailed.Error String) ( Http.Metadata, String ))
    | GotFlowData (Result (Http.Detailed.Error String) ( Http.Metadata, String ))
    | UpdateMasking String
    | UpdateInput String
    | UpdateMaskingAndInput Sandbox
    | Refresh
    | Error String
    | ChangeMaskingView MaskingView
    | ClosePopup


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
