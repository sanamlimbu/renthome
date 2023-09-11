import { Box } from "@mui/material";
import { Outlet } from "react-router-dom";
import Footer from "../components/footer";
import NavBar from "../components/navBar";

export default function RootPage() {
  // const matchesXs = useMediaQuery("(max-width: 599px)");
  // const matchesSm = useMediaQuery("(min-width: 600px) and (max-width: 959px)");
  // const matchesMd = useMediaQuery("(min-width: 960px) and (max-width: 1279px)");
  // const matchesLg = useMediaQuery("(min-width: 1280px)");

  // const padding = matchesXs
  //   ? "2vh 5vw"
  //   : matchesSm
  //   ? "3vh 6vw"
  //   : matchesMd
  //   ? "4vh 7vw"
  //   : matchesLg
  //   ? "5vh 22vw"
  //   : "5vh 30vw";

  return (
    <>
      <NavBar />
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          background: "rgb(246,245,247)",
        }}
      >
        <Outlet />
      </Box>
      <Footer />
    </>
  );
}
