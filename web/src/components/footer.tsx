import styled from "@emotion/styled";
import {
  FacebookOutlined,
  LinkedIn,
  Pinterest,
  Twitter,
  YouTube,
} from "@mui/icons-material";
import { Box, Icon, IconProps, Typography } from "@mui/material";
import { blue } from "@mui/material/colors";

export default function Footer() {
  return (
    <Box
      sx={{
        paddingLeft: "10rem",
        paddingRight: "10rem",
        marginBottom: "2em",
        marginTop: "3em",
      }}
    >
      <Box
        sx={{
          display: "flex",
          justifyContent: "space-between",
          flexWrap: "wrap",
          gap: "10px",
        }}
      >
        <Box sx={{ display: "flex", flexWrap: "wrap", gap: "12px" }}>
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
        </Box>
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
      </Box>
      <Typography
        sx={{
          fontSize: "14px",
          opacity: "0.5",
          textAlign: "center",
          marginTop: "1em",
          fontWeight: "bold",
        }}
      >
        Copyright © {new Date().getFullYear()} renthome.com.au, all rights
        reserved.
      </Typography>
    </Box>
  );
}

const LinkSpan = styled("span")({
  fontSize: "14px",
  opacity: "0.5",
  fontWeight: "bold",
  "&:hover": {
    textDecoration: "underline",
    color: blue[500],
    opacity: "1",
  },
  cursor: "pointer",
});

const SocialIcon = styled(Icon)<IconProps>(() => ({
  opacity: "0.5",
  "&:hover": {
    opacity: "1",
  },
  cursor: "pointer",
}));
