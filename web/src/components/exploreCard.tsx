import { Box, Card, Link, Typography } from "@mui/material";
export interface ExploreCardProps {
  title: string;
  shortDescription: string;
  linkText: string;
  imgSrc: string;
}
export default function ExploreCard({
  title,
  shortDescription,
  linkText,
  imgSrc,
}: ExploreCardProps) {
  return (
    <Card
      variant="outlined"
      sx={{ maxWidth: "24em", borderRadius: "12px", cursor: "pointer" }}
    >
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          margin: "1em",
        }}
      >
        <img src={imgSrc} height={98} />
      </Box>
      <Box sx={{ margin: "0 1em 1em 1em" }}>
        <Typography variant="h4">{title}</Typography>
        <Typography
          sx={{
            color: "dimgray",
            marginTop: "0.8em",
            marginBottom: "0.8em",
          }}
        >
          {shortDescription}
        </Typography>
        <Link
          sx={{
            fontWeight: "600",
            cursor: "pointer",
            textDecoration: "none",
            "&:hover": {
              textDecoration: "underline",
            },
          }}
        >
          {linkText}
        </Link>
      </Box>
    </Card>
  );
}
