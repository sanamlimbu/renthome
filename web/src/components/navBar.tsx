import { Close, Menu } from "@mui/icons-material";
import {
  Button,
  ButtonProps,
  Divider,
  Typography,
  useMediaQuery,
} from "@mui/material";
import { styled } from "@mui/material/styles";
import React from "react";
import { useNavigate } from "react-router-dom";
import RentHomeLogo from "./rentHomeLogo";

export default function NavBar() {
  const [openMenu, setOpenMenu] = React.useState(false);
  const matches = useMediaQuery("(min-width:769px)");
  const navigate = useNavigate();

  return (
    <>
      <div
        style={{
          display: "flex",
          justifyContent: matches ? "space-evenly" : "space-between",
          alignItems: "center",
          padding: matches ? "20px 0 20px 0" : "20px 10px 20px 10px",
        }}
      >
        <div style={{ display: "flex" }}>
          <Typography
            sx={{
              display: "flex",
              alignItems: "center",
              gap: "6px",
              fontWeight: "600",
              padding: "10px",
              borderRadius: "4px",
              cursor: "pointer",
              "&:hover": {
                background: "lightgray",
              },
            }}
            onClick={() => setOpenMenu((prevState) => !prevState)}
          >
            {openMenu ? (
              matches ? (
                <Close />
              ) : (
                <Close fontSize="large" />
              )
            ) : matches ? (
              <Menu />
            ) : (
              <Menu fontSize="large" />
            )}
            {matches && "Menu"}
          </Typography>
          {!matches && <RentHomeLogo />}
        </div>
        {matches && <RentHomeLogo />}
        <div>
          <NavBarButton
            sx={{
              marginRight: "10px",
            }}
            onClick={() => navigate("/login")}
          >
            Sign in
          </NavBarButton>
          <NavBarButton variant="contained" onClick={() => navigate("/signup")}>
            Join
          </NavBarButton>
        </div>
      </div>
      <Divider />
      {openMenu && (
        <div>
          <div
            style={{
              display: "flex",
              justifyContent: "center",
              paddingTop: "20px",
              paddingBottom: "20px",
              gap: "20px",
            }}
          >
            <NavBarButton>Buy</NavBarButton>
            <NavBarButton>Rent</NavBarButton>
            <NavBarButton>Sold</NavBarButton>
            <NavBarButton>Find agents</NavBarButton>
          </div>
          <div style={{ backgroundColor: "black", width: "100%" }}></div>
          <Divider />
        </div>
      )}
    </>
  );
}

const NavBarButton = styled(Button)<ButtonProps>(({ theme }) => ({
  textTransform: "none",
  fontWeight: "600",
}));
