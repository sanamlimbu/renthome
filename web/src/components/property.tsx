import {
  BedOutlined,
  DirectionsCarOutlined,
  MoreVert,
  ShowerOutlined,
  StarOutline,
} from "@mui/icons-material";
import {
  Avatar,
  Card,
  CardContent,
  CardHeader,
  Divider,
  IconButton,
  Typography,
} from "@mui/material";
import { red } from "@mui/material/colors";
import { Carousel } from "react-responsive-carousel";
import "react-responsive-carousel/lib/styles/carousel.min.css";
import { useNavigate } from "react-router-dom";
import imageOne from "../assets/1.jpg";
import imageTwo from "../assets/2.jpg";
import imageThree from "../assets/3.jpg";
import imageFour from "../assets/4.jpg";

export default function Property() {
  const navigate = useNavigate();

  return (
    <Card
      variant="outlined"
      sx={{ maxWidth: "40rem", cursor: "pointer" }}
      onClick={() => navigate("/")}
    >
      <CardHeader
        avatar={<Typography fontWeight={"bold"}>Perth Realty</Typography>}
        action={
          <div style={{ display: "flex", alignItems: "center", gap: "8px" }}>
            <Typography>Sanam Limbu</Typography>
            <Avatar sx={{ bgcolor: red[500] }} aria-label="recipe">
              R
            </Avatar>
          </div>
        }
        sx={{ backgroundColor: "lightgreen" }}
      />
      <Carousel showArrows={true} showThumbs={false}>
        <div>
          <img src={imageOne} />
        </div>
        <div>
          <img src={imageTwo} />
        </div>
        <div>
          <img src={imageThree} />
        </div>
        <div>
          <img src={imageFour} />
        </div>
      </Carousel>
      <CardContent>
        <div style={{ display: "flex", justifyContent: "space-between" }}>
          <Typography sx={{ fontWeight: "bold" }}>{"$350"}</Typography>
          <div>
            <IconButton>
              <StarOutline />
            </IconButton>
            <IconButton>
              <MoreVert />
            </IconButton>
          </div>
        </div>
        <Typography sx={{ marginBottom: "10px" }}>
          {"50/12 Wall Street, Maylands"}
        </Typography>
        <div style={{ display: "flex", gap: "12px", height: "20px" }}>
          <Typography
            sx={{ display: "flex", alignItems: "center", gap: "6px" }}
          >
            <BedOutlined />
            <span>2</span>
          </Typography>
          <Typography
            sx={{ display: "flex", alignItems: "center", gap: "6px" }}
          >
            <ShowerOutlined />
            <span>1</span>
          </Typography>
          <Typography
            sx={{ display: "flex", alignItems: "center", gap: "6px" }}
          >
            <DirectionsCarOutlined />
            <span>1</span>
          </Typography>
          <Divider orientation="vertical" />
          <Typography>Unit</Typography>
        </div>
      </CardContent>
    </Card>
  );
}
