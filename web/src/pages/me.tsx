import { EmailOutlined, LockOutlined } from "@mui/icons-material";
import {
  Alert,
  Box,
  Button,
  Divider,
  Snackbar,
  Typography,
} from "@mui/material";
import { red } from "@mui/material/colors";
import React, { useContext } from "react";
import { useNavigate } from "react-router-dom";
import { API_ADDRESS } from "../config";
import { UserContext } from "../context/user";
import {
  getTokenFromLocalStorage,
  removeTokenFromLocalStorage,
  removeUserFromLocalStorage,
} from "../helpers/auth";

export function MePage() {
  const { user } = useContext(UserContext);
  const navigate = useNavigate();
  const [error, setError] = React.useState("");
  const [openSnackbar, setOpenSnackbar] = React.useState(false);
  const { setUser } = useContext(UserContext);
  // if (user === undefined) {
  //   navigate("login");
  // }

  const handleSnackbarClose = () => {
    setOpenSnackbar(false);
  };

  const handleLogout = async () => {
    try {
      const res = await fetch(`${API_ADDRESS}/auth/logout`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${getTokenFromLocalStorage()}`,
        },
        body: JSON.stringify({ user_id: `${user?.id}` }),
      });
      if (!res.ok) {
        const data = await res.json();
        setError(data.message);
        setOpenSnackbar(true);
      } else {
        removeUserFromLocalStorage();
        removeTokenFromLocalStorage();
        setUser(undefined);
        setOpenSnackbar(false);
      }
    } catch (error) {
      setOpenSnackbar(true);
    }
  };
  return (
    <Box sx={{ display: "flex", flexDirection: "column", gap: "24px" }}>
      <div style={{ display: "flex", justifyContent: "space-between" }}>
        <Typography sx={{ fontWeight: "700", fontSize: "28px" }}>
          Account overview
        </Typography>
        <Button
          variant="contained"
          sx={{ textTransform: "none", fontWeight: "700" }}
          onClick={handleLogout}
        >
          Log out
        </Button>
      </div>
      <div>
        <Typography fontWeight={"600"} paddingBottom="10px">
          Email
        </Typography>
        <div style={{ display: "flex", justifyContent: "space-between" }}>
          <div>
            <div style={{ display: "flex", alignItems: "center", gap: "8px" }}>
              <EmailOutlined />
              <Typography>{"sudosanam@gmail.com"}</Typography>
            </div>
          </div>
          <Button
            sx={{ textTransform: "none", fontWeight: "bold" }}
            onClick={() => navigate("/email-update")}
          >
            Update
          </Button>
        </div>
        <Typography
          sx={{
            color: "rgb(92,158,141)",
            paddingLeft: "26px",
          }}
        >
          {"Verified"}
        </Typography>
      </div>
      <div>
        <Typography fontWeight={"600"} paddingBottom="10px">
          Password
        </Typography>
        <div style={{ display: "flex", justifyContent: "space-between" }}>
          <div style={{ display: "flex", alignItems: "center", gap: "8px" }}>
            <LockOutlined />
            <Typography>{"********"}</Typography>
          </div>
          <Button
            sx={{ textTransform: "none", fontWeight: "bold" }}
            onClick={() => navigate("/password-update")}
          >
            Update
          </Button>
        </div>
      </div>
      <div>
        <Typography fontWeight={"600"}>{"Sign out on all devices"}</Typography>
        <Typography sx={{ paddingTop: "20px", paddingBottom: "20px" }}>
          {
            "Lost a device or recently used a public computer? Protect your account by signing out on all devices."
          }
        </Typography>
        <div>
          <Button
            variant="outlined"
            sx={{ textTransform: "none", fontWeight: "bold" }}
          >
            {" "}
            Sign out on all devices
          </Button>
        </div>
      </div>
      <div>
        <Button
          sx={{ color: red[500], textTransform: "none", fontWeight: "bold" }}
        >
          Delete account
        </Button>
      </div>
      <Divider />
      <Typography sx={{ fontWeight: "700", fontSize: "28px" }}>
        Notification settings
      </Typography>
      <Typography fontWeight={"600"}>Property journey</Typography>
      <Typography fontWeight={"600"}>Properties</Typography>
      <Typography fontWeight={"600"}>Property market</Typography>
      <Typography fontWeight={"600"}>Finance</Typography>
      <Divider />
      <Typography sx={{ fontWeight: "700", fontSize: "28px" }}>
        Data privacy
      </Typography>
      <div></div>
      <Snackbar
        open={openSnackbar}
        autoHideDuration={6000}
        onClose={handleSnackbarClose}
        anchorOrigin={{
          vertical: "bottom",
          horizontal: "center",
        }}
      >
        <Alert
          onClose={handleSnackbarClose}
          severity="error"
          sx={{ width: "100%" }}
        >
          {error || "Something went wrong, please try again."}
        </Alert>
      </Snackbar>
    </Box>
  );
}
