import { Button, Card, Typography } from "@mui/material";
import { red } from "@mui/material/colors";
import { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";
import RentHomeLogo from "../../components/rentHomeLogo";
import { API_ADDRESS } from "../../config";
import { UserContext } from "../../context/user";
import {
  getTokenFromLocalStorage,
  removeTokenFromLocalStorage,
} from "../../helpers/auth";
import "../../styles/index.css";

export default function SignoutAllPage() {
  const [successful, setSuccessful] = useState(false);
  const [error, setError] = useState("");
  const navigate = useNavigate();
  const { setUser } = useContext(UserContext);

  const handleSignoutAllDeivces = async () => {
    try {
      const res = await fetch(`${API_ADDRESS}/api/auth/signout-all`, {
        method: "POST",
        headers: {
          Authorization: `Bearer ${getTokenFromLocalStorage()}`,
          "X-User-Agent": window.navigator.userAgent,
        },
      });
      if (res.ok) {
        setSuccessful(true);
      } else {
        const data = await res.json();
        setError(data.message);
      }
    } catch (error) {
      setError("Something went wrong try again.");
    }
  };

  return (
    <div className="Container-centered">
      <Card variant="outlined" className="Card">
        <RentHomeLogo />
        {successful ? (
          <div
            style={{ display: "flex", flexDirection: "column", gap: "12px" }}
          >
            <Typography
              textAlign={"center"}
              fontWeight={"700"}
              fontSize={"24px"}
            >
              All of your devices have been logged out
            </Typography>
            <Typography
              textAlign={"center"}
              paddingTop="10px"
              paddingBottom="10px"
            >
              It may take up to 15 minutes before all devices are logged out.
            </Typography>

            <div style={{ padding: "0 30px 0 30px" }}>
              <Button
                variant="contained"
                sx={{ textTransform: "none", fontWeight: "700" }}
                fullWidth
                onClick={() => {
                  removeTokenFromLocalStorage();
                  setUser(undefined);
                }}
              >
                Done
              </Button>
            </div>
          </div>
        ) : (
          <div
            style={{ display: "flex", flexDirection: "column", gap: "12px" }}
          >
            <Typography
              textAlign={"center"}
              fontWeight={"700"}
              fontSize={"24px"}
            >
              Logout of all your devices
            </Typography>
            <Typography textAlign={"center"} padding="0 30px 0 30px">
              All your devices will be logged out across all sites. You will no
              longer receive notifications for any devices previously logged
              into
            </Typography>
            {error && (
              <Typography
                color={red[500]}
                textAlign="center"
                paddingTop={"10px"}
                paddingBottom={"10px"}
              >
                {error}
              </Typography>
            )}
            <div style={{ padding: "0 30px 0 30px" }}>
              <Button
                variant="contained"
                size="large"
                sx={{ textTransform: "none", fontWeight: "600" }}
                type="submit"
                fullWidth
                onClick={handleSignoutAllDeivces}
              >
                Logout
              </Button>
            </div>

            <div style={{ padding: "0 30px 0 30px" }}>
              <Button
                sx={{ textTransform: "none", fontWeight: "700" }}
                fullWidth
                onClick={() => navigate("/me")}
              >
                Cancel
              </Button>
            </div>
          </div>
        )}
      </Card>
    </div>
  );
}
