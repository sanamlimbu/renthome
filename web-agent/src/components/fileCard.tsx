import { CancelRounded } from "@mui/icons-material";
import { Box, Card, CardMedia, IconButton, Typography } from "@mui/material";

interface ImageCardProps {
  src: string;
  name: string;
  remove: (name: string) => void;
}
export default function FileCard({ src, name, remove }: ImageCardProps) {
  const handleRemove = () => {
    remove(name);
  };
  return (
    <Box>
      <Card sx={{ position: "relative" }}>
        <CardMedia sx={{ height: 100, width: 100 }} image={src} />
        <IconButton
          sx={{ color: "red", position: "absolute", top: 0, right: 0 }}
          onClick={handleRemove}
        >
          <CancelRounded />
        </IconButton>
      </Card>
      <Typography>{name}</Typography>
    </Box>
  );
}
