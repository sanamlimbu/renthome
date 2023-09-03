import {
  Action,
  Client,
  ClientContextProvider,
  createClient,
} from "react-fetching-library";
import { AuthContextProvider } from "./context/authContext";
import RoutesProvider from "./routes/route";

const API_ENDPOINT = import.meta.env.VITE_RENTHOME_API_ENDPOINT;

console.log(API_ENDPOINT);

const prefixURL =
  (prefix: string) => (client: Client) => async (action: Action) => {
    return {
      ...action,
      endpoint: `${window.location.protocol}//${API_ENDPOINT}${prefix}${action.endpoint}`,
    };
  };

const client = createClient({
  requestInterceptors: [prefixURL("/api")],
});

function App() {
  return (
    <ClientContextProvider client={client}>
      <AuthContextProvider>
        <RoutesProvider />
      </AuthContextProvider>
    </ClientContextProvider>
  );
}

export default App;
