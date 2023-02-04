import { Apple, Facebook, Google } from "@mui/icons-material";
import { styled, Typography } from "@mui/material";

enum SocialAction {
  signin = "SIGNIN",
  signup = "SIGNUP",
}

interface SocialType {
  name: string;
  bgColor: string;
  color: string;
}

const socialList: SocialType[] = [
  {
    name: "Google",
    bgColor: "lightgray",
    color: "black",
  },
  {
    name: "Facebook",
    bgColor: "white",
    color: "black",
  },
  {
    name: "Apple",
    bgColor: "black",
    color: "white",
  },
];

interface ISocial {
  type: SocialType;
  action: SocialAction;
}

export default function Social(props: ISocial) {
  const { type, action } = props;
  const handleClick = () => console.log(action);

  const StyledDiv = styled("div")({
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
  });

  return (
    <StyledDiv onClick={handleClick}>
      {type.name === "Google" && <Google />}
      {type.name === "Facebook" && <Facebook />}
      {type.name === "Apple" && <Apple />}
      <Typography
        sx={{ flexGrow: "1", textAlign: "center", fontWeight: "600" }}
      >
        Continue with {type.name}
      </Typography>
    </StyledDiv>
  );
}

export { SocialAction, socialList };
