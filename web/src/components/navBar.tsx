import { AccountBox, Close, Menu } from "@mui/icons-material";
import {
  Button,
  ButtonProps,
  Divider,
  IconButton,
  Typography,
  useMediaQuery,
} from "@mui/material";
import { styled } from "@mui/material/styles";
import React, { useContext } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { UserContext } from "../context/user";
import RentHomeLogo from "./rentHomeLogo";

export default function NavBar() {
  const [openMenu, setOpenMenu] = React.useState(false);
  const matches = useMediaQuery("(min-width:769px)");
  const navigate = useNavigate();
  const { user } = useContext(UserContext);
  const location = useLocation();
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
          {(!matches || location.pathname === "/me") && (
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
          )}
          {!matches && <RentHomeLogo />}
        </div>
        {matches && <RentHomeLogo />}
        {location.pathname !== "/me" && matches && (
          <div>
            <NavBarButton onClick={() => navigate("/rent")}>Rent</NavBarButton>
            <NavBarButton onClick={() => navigate("/buy")}>Buy</NavBarButton>
            <NavBarButton onClick={() => navigate("/sold")}>Sold</NavBarButton>
            <NavBarButton onClick={() => navigate("/find-agent")}>
              Find agents
            </NavBarButton>
          </div>
        )}
        {user ? (
          <div>
            <IconButton
              style={{
                borderRadius: "4px",
                padding: "8px",
              }}
              onClick={() => navigate("/me")}
            >
              <AccountBox />
            </IconButton>
          </div>
        ) : (
          <div>
            <NavBarButton
              sx={{
                marginRight: "10px",
              }}
              onClick={() => navigate("/login")}
            >
              Sign in
            </NavBarButton>
            <NavBarButton
              variant="contained"
              onClick={() => navigate("/signup")}
            >
              Join
            </NavBarButton>
          </div>
        )}
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
            <NavBarButton onClick={() => navigate("/buy")}>Buy</NavBarButton>
            <NavBarButton onClick={() => navigate("/rent")}>Rent</NavBarButton>
            <NavBarButton onClick={() => navigate("/sold")}>Sold</NavBarButton>
            <NavBarButton onClick={() => navigate("/find-agent")}>
              Find agents
            </NavBarButton>
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
