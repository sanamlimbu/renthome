import { Home } from "@mui/icons-material";
import { Typography } from "@mui/material";
import { useNavigate } from "react-router-dom";
import "../styles/index.css";

export default function RentHomeLogo() {
  const navigate = useNavigate();
  return (
    <Typography
      className="Typography-centered"
      onClick={() => navigate("/")}
      sx={{ cursor: "pointer" }}
    >
      <Home color="primary" fontSize="large" />
      <span style={{ fontWeight: "700", fontSize: "22px" }}>renthome.com</span>
    </Typography>
  );
}
