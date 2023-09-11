import { ThemeProvider } from "@emotion/react";
import { ClientContextProvider, createClient } from "react-fetching-library";
import { UserContextProvider } from "./context/user";
import RoutesProvider from "./routes/route";
import theme from "./theme";

const client = createClient();

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
