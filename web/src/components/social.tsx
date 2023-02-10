import { Apple, Facebook, Google } from "@mui/icons-material";
import { styled, Typography } from "@mui/material";

import { GOOGLE_CLIENT_ID, GOOGLE_REDIRECT_URI } from "../config";
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
  from: string;
}

export default function Social(props: ISocial) {
  const { type, action, from } = props;
  const url = getSocialURL(type, action, from);

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

function getSocialURL(
  type: SocialType,
  action: SocialAction,
  from: string
): string {
  let url = "";
  switch (type.name) {
    case "Google": {
      const rootUrl = `https://accounts.google.com/o/oauth2/v2/auth`;
      const options = {
        redirect_uri: GOOGLE_REDIRECT_URI as string,
        client_id: GOOGLE_CLIENT_ID as string,
        access_type: "offline",
        response_type: "code",
        prompt: "consent",
        scope: [
          "https://www.googleapis.com/auth/userinfo.profile",
          "https://www.googleapis.com/auth/userinfo.email",
        ].join(" "),
        state: from,
      };

      if (action === SocialAction.signup) {
        options.redirect_uri = process.env
          .RENTHOME_GOOGLE_REDIRECT_URL_SIGNUP as string;
      }

      const qs = new URLSearchParams(options);
      url = `${rootUrl}?${qs.toString()}`;
      break;
    }
    case "Facebook": {
      break;
    }
    case "Apple": {
      break;
    }
  }

  return url;
}

export { SocialAction, socialList };
