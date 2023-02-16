import { Apple, Facebook, Google } from "@mui/icons-material";
import { styled, Typography } from "@mui/material";
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
      {type.name === "Google" && <Google />}
      {type.name === "Facebook" && <Facebook />}
      {type.name === "Apple" && <Apple />}
      <Typography
        sx={{ flexGrow: "1", textAlign: "center", fontWeight: "600" }}
      >
        Continue with {type.name}
      </Typography>
    </StyledAnchor>
  );
}
