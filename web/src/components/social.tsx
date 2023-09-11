import { styled, Typography } from "@mui/material";
import Facebook from "../assets/icons/facebook.svg";
import Google from "../assets/icons/google.svg";
import { getSocialURL } from "../helpers/auth";
import { SocialType } from "../types/types";

export default function Social(props: { type: SocialType }) {
  const { type } = props;
  const url = getSocialURL(type);

  const StyledAnchor = styled("a")({
    display: "flex",
    border: "solid 1px lightgray",
    borderRadius: "15px",
    padding: "10px",
    cursor: "pointer",
    flexWrap: "wrap",
    color: type.color,
    backgroundColor: type.bgColor,
    "&:hover": {
      opacity: "0.8",
    },
    textDecoration: "none",
  });

  return (
    <StyledAnchor href={url}>
      {type.name === "Google" && <img src={Google} height={"24px"} />}
      {type.name === "Facebook" && <img src={Facebook} height={"24px"} />}
      <Typography
        sx={{ flexGrow: "1", textAlign: "center", fontWeight: "600" }}
      >
        Continue with {type.name}
      </Typography>
    </StyledAnchor>
  );
}
