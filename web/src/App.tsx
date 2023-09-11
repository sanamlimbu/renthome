import { ThemeProvider } from "@emotion/react";
import {
  Action,
  Client,
  ClientContextProvider,
  createClient,
} from "react-fetching-library";
import { API_ADDRESS } from "./config";
import { UserContextProvider } from "./context/user";
import RoutesProvider from "./routes/route";
import theme from "./theme";

const prefixURL =
  (prefix: string) => (_client: Client) => async (action: Action) => {
    return {
      ...action,
      endpoint: `${window.location.protocol}//${API_ADDRESS}${prefix}${action.endpoint}`,
    };
  };

const client = createClient({
  requestInterceptors: [prefixURL("/api")],
});

function App() {
  return (
    <ClientContextProvider client={client}>
      <ThemeProvider theme={theme}>
        <UserContextProvider>
          <RoutesProvider />
        </UserContextProvider>
      </ThemeProvider>
    </ClientContextProvider>
  );
}

export default App;
