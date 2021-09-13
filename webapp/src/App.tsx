import React, { createContext } from "react";
import { BrowserRouter, Route } from "react-router-dom";
import { Navbar } from "./components/navbar";
import { HomePage } from "./pages/HomePage";
import { SignInPage } from "./pages/SignInPage";
import { ApiClient } from "./api/client";
import { SeisekiPage } from "./pages/SeisekiPage";
import { CookiesProvider } from "react-cookie";
import { RegisterPage } from "./pages/RegisterPage";

export const ApiClientContext = createContext(
  new ApiClient("http://localhost:8081/api")
);

function App(): JSX.Element {
  return (
    <CookiesProvider>
      <ApiClientContext.Provider
        value={new ApiClient("http://localhost:8081/api")}
      >
        <BrowserRouter>
          <Navbar />
          <Route exact path="/" component={HomePage} />
          <Route exact path="/auth/signin" component={SignInPage} />
          <Route exact path="/auth/register" component={RegisterPage} />
          <Route exact path="/seiseki" component={SeisekiPage} />
        </BrowserRouter>
      </ApiClientContext.Provider>
    </CookiesProvider>
  );
}

export default App;
