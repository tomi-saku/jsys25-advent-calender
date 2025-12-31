import { GoogleLogin, type CredentialResponse } from "@react-oauth/google";
import { type UserData } from "../App";

type GoogleLoginButtonProps = {
  handleValueChange: (data: UserData) => void;
};

const GoogleLoginButton = ({ handleValueChange }: GoogleLoginButtonProps) => {
  const handleLoginSuccess = (credentialResponse: CredentialResponse) => {
    //console.log(credentialResponse);
    fetch("http://localhost:8080/authorization", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${credentialResponse.credential}`,
      },
    })
      .then((response) => response.json())
      .then((data: UserData) => {
        console.log("Email adress: ", data.email);
        console.log("Image: ", data.image);
        handleValueChange(data);
      })
      .catch((error) => {
        console.log("Error: ", error);
      });
  };

  const handleLoginError = () => {
    console.log("Login Failed");
  };

  return (
    <GoogleLogin onSuccess={handleLoginSuccess} onError={handleLoginError} />
  );
};

export default GoogleLoginButton;
