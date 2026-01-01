import { GoogleLogin, type CredentialResponse } from "@react-oauth/google";
import { type UserData } from "../App";

type GoogleLoginButtonProps = {
  handleValueChange: (data: UserData) => void;
};

const GoogleLoginButton = ({ handleValueChange }: GoogleLoginButtonProps) => {
  // ログイン成功時の処理
  const handleLoginSuccess = (credentialResponse: CredentialResponse) => {
    // バックエンドへのリクエスト
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
        //親要素へユーザーデータを渡す
        handleValueChange(data);
      })
      .catch((error) => {
        console.log("Error: ", error);
      });
  };
  // ログイン失敗時の処理
  const handleLoginError = () => {
    console.log("Login Failed");
  };

  return (
    <GoogleLogin onSuccess={handleLoginSuccess} onError={handleLoginError} />
  );
};

export default GoogleLoginButton;
