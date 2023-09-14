import { AddCircle, LocationOn } from "@mui/icons-material";
import { Box, Typography } from "@mui/material";

export default function SuggestedLocationCard(props: {
  location: string;
  handleSelect: (value: string) => void;
}) {
  const { location, handleSelect } = props;
  return (
    <Box
      sx={{
        padding: "0.5em",
        display: "flex",
        alignItems: "center",
        "&:hover": {
          background: "rgb(246,245,247)",
        },
        cursor: "pointer",
      }}
      onClick={() => handleSelect(location)}
    >
      <Box
        sx={{
          background: "rgb(229,227,232)",
          padding: "6px",
          borderRadius: "4px",
        }}
      >
        <LocationOn fontSize="small" />
      </Box>
      <Typography
        sx={{
          marginLeft: "1em",
          marginRight: "1em",
        }}
      >
        {location}
      </Typography>
      <Box
        sx={{
          flexGrow: "1",
          display: "flex",
          justifyContent: "flex-end",
        }}
      >
        <AddCircle fontSize="small" />
      </Box>
    </Box>
  );
}
