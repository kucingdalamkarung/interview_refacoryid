import React, { useState } from "react";
import "./App.css";
import GitHubLogin from "./GithubLogin";

function App() {
  const [name, setName] = useState("");

  return (
    <div className="App">
      <header className="App-header">
        {name && <h2>Hai {name}</h2>}
        {/* add client id, // add clientSecret */}
        <GitHubLogin
          clientId="a064bc8979a2f7ba1a6d"
          clientSecret="f77d6d53c3ab2ac54cb9502f21345622ececf757"
          redirectUri="http://localhost:3000"
          onSuccess={(name) => setName(name)}
          onFailure={(resp) => console.log(resp)}
        />
      </header>
    </div>
  );
}

export default App;
