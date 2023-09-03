import { Avatar, Box, Button, Card, Typography } from "@mui/material";
export interface MortgageBrokersCardProps {
  imgSrc: string;
  companyLogoSrc: string;
  companyName: string;
  name: string;
  locationCity: string;
}
export default function MortgageBrokersCard({
  imgSrc,
  companyLogoSrc,
  companyName,
  name,
  locationCity,
}: MortgageBrokersCardProps) {
  return (
    <Card variant="outlined" sx={{ maxWidth: "24em", borderRadius: "12px" }}>
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          flexDirection: "column",
        }}
      >
        <Box
          sx={{
            backgroundColor: "#006775",
            padding: "1em",
          }}
        >
          <img src={companyLogoSrc} />
          <Typography fontSize={"14px"}> {companyName}</Typography>
        </Box>
        <Box
          sx={{
            paddingLeft: "1.2em",
            paddingTop: "1.2em",
            paddingRight: "1.2em",
          }}
        >
          <Avatar />
          <Typography
            sx={{
              fontWeight: "bold",
            }}
          >
            {name}
          </Typography>
          <Typography>Mortgage broker</Typography>
          <Typography
            marginTop={"0.5em"}
            marginBottom={"0.5em"}
            fontSize={"15px"}
          >
            Based in {locationCity}
          </Typography>
          <Typography
            sx={{
              marginTop: "2em",
              "&:hover": {
                textDecoration: "underline",
              },
            }}
            fontSize={"15px"}
          >
            See more details
          </Typography>
        </Box>
        <Box
          sx={{
            padding: "1em",
          }}
        >
          <Button
            variant="outlined"
            sx={{
              textTransform: "none",
            }}
            color="inherit"
          >
            Request a call back
          </Button>
        </Box>
      </Box>
    </Card>
  );
}
