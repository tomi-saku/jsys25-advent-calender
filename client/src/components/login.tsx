import { GoogleLogin, type CredentialResponse } from "@react-oauth/google";

type User = {
  email: string;
  image: string;
};

export default function GoogleLoginButton() {
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
      .then((data: User) => {
        console.log("Email adress: ", data.email);
        console.log("Image: ", data.image);
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
}
