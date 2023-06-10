import { ThemeProvider } from "@emotion/react";
import { UserContextProvider } from "./context/user";
import RoutesProvider from "./routes/route";
import theme from "./theme";

function App() {
  return (
    <ThemeProvider theme={theme}>
      <UserContextProvider>
        <RoutesProvider />
      </UserContextProvider>
    </ThemeProvider>
  );
}

export default App;
