import styled from "@emotion/styled";
import {
  FacebookOutlined,
  LinkedIn,
  Pinterest,
  Twitter,
  YouTube,
} from "@mui/icons-material";
import { Icon, IconProps, Typography } from "@mui/material";
import { blue } from "@mui/material/colors";

export default function Footer() {
  return (
    <div
      style={{
        paddingLeft: "20vw",
        paddingRight: "20vw",
        marginBottom: "2em",
        marginTop: "3em",
      }}
    >
      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          flexWrap: "wrap",
          gap: "10px",
        }}
      >
        <div style={{ display: "flex", flexWrap: "wrap", gap: "8px" }}>
          <SocialIcon>
            <FacebookOutlined />
          </SocialIcon>
          <SocialIcon>
            <Twitter />
          </SocialIcon>
          <SocialIcon>
            <Pinterest />
          </SocialIcon>
          <SocialIcon>
            <LinkedIn />
          </SocialIcon>
          <SocialIcon>
            <YouTube />
          </SocialIcon>
        </div>
        <Typography sx={{ display: "flex", gap: "1em" }}>
          <LinkSpan>Advertise with us</LinkSpan>
          <LinkSpan>Contact us</LinkSpan>
          <LinkSpan>Agent admin</LinkSpan>
          <LinkSpan> Media sales</LinkSpan>
          <LinkSpan>Legal</LinkSpan>
          <LinkSpan>Privacy</LinkSpan>
          <LinkSpan>Site map </LinkSpan>
          <LinkSpan>Career</LinkSpan>
        </Typography>
      </div>
      <Typography
        sx={{
          fontSize: "14px",
          opacity: "0.5",
          textAlign: "center",
          marginTop: "1em",
        }}
      >
        Copyright © {new Date().getFullYear()} renthome.com.au, all rights
        reserved.
      </Typography>
    </div>
  );
}

const LinkSpan = styled("span")({
  fontSize: "14px",
  opacity: "0.5",
  "&:hover": {
    textDecoration: "underline",
    color: blue[500],
    opacity: "1",
  },
  cursor: "pointer",
});

const SocialIcon = styled(Icon)<IconProps>(({ theme }) => ({
  opacity: "0.5",
  "&:hover": {
    opacity: "1",
  },
  cursor: "pointer",
  size: "small",
}));
