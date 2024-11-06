import { ChangeEvent } from "react";
import { CredentialResponse as GoogleCredentialResponse } from "google-one-tap";

declare global {
  namespace Google {
    type CredentialResponse = GoogleCredentialResponse;
  }
  type InputChangeEvent = ChangeEvent<HTMLInputElement>;
}

export {};
