import { useState } from "react";
import "./App.css";
import GoogleLoginButton from "./components/login";

export type UserData = {
  email: string;
  image: string;
};

function App() {
  const [userData, setUserData] = useState<UserData | null>(null);

  const getUserData = (newData: UserData) => {
    setUserData(newData);
  };

  if (userData) {
    return (
      <div>
        <img src={userData.image} alt="User profile" />
        <p>ようこそ、{userData.email}さん</p>
      </div>
    );
  } else {
    return (
      <>
        <p>ログインテスト</p>
        <GoogleLoginButton handleValueChange={getUserData} />
      </>
    );
  }
}

export default App;
