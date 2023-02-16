import { Box, useMediaQuery } from "@mui/material";
import { useContext } from "react";
import { Outlet } from "react-router-dom";
import Footer from "../components/footer";
import NavBar from "../components/navBar";
import { UserContext } from "../context/user";

export default function RootPage() {
  const { user } = useContext(UserContext);
  const matches = useMediaQuery("(min-width:769px)");

  return (
    <>
      <NavBar />
      <Box sx={{ padding: matches ? "5vh 20vw 5vh 20vw" : "2vh 5vw 2vh 5vw" }}>
        <Outlet />
      </Box>
      <Footer />
    </>
  );
}
