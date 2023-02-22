import { UserContextProvider } from "./context/user";
import RoutesProvider from "./routes/route";

function App() {
  return (
    <UserContextProvider>
      <RoutesProvider />
    </UserContextProvider>
  );
}

export default App;
