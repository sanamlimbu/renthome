import { Close } from "@mui/icons-material";
import { Box, IconButton, Typography } from "@mui/material";

export default function SelelectedLocationCard(props: {
  location: string;
  handleCancel: (value: string) => void;
}) {
  const { location, handleCancel } = props;
  return (
    <Box
      sx={{
        display: "flex",
        alignItems: "center",
        background: "rgb(229,227,232)",
        padding: "4px 4px 4px 6px",
        borderRadius: "16px",
      }}
    >
      <Typography fontWeight="600" fontSize={"15px"}>
        {location}
      </Typography>
      <IconButton
        onClick={() => handleCancel(location)}
        size="small"
        sx={{
          padding: 0,
          margin: 0,
        }}
      >
        <Close fontSize="small" />
      </IconButton>
    </Box>
  );
}
