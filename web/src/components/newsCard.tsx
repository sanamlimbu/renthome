import { Box, Typography } from "@mui/material";

export interface NewsCardProps {
  title: string;
  readTime: string;
  imgSrc: string;
}
export default function NewsCard({ title, readTime, imgSrc }: NewsCardProps) {
  return (
    <Box sx={{ cursor: "pointer" }}>
      <img src={imgSrc} height="160px" />
      <Typography
        sx={{
          maxWidth: "240px",
          overflow: "hidden",
          textOverflow: "ellipsis",
          fontWeight: "500",
        }}
      >
        {title}
      </Typography>
      <Typography marginTop={"0.5em"}>{readTime}</Typography>
    </Box>
  );
}
