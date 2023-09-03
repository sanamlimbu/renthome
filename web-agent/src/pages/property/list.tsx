import { Box, CircularProgress, Typography } from "@mui/material";
import { useQuery } from "react-fetching-library";
import { fetching } from "../../fetching";

export default function PropertyListPage() {
  const {
    loading,
    payload: properties,
    error,
  } = useQuery(fetching.query.getPropertyList);

  return (
    <Box>
      {loading && <CircularProgress />}
      {error && <Typography>{error.valueOf.name}</Typography>}

      {properties?.map((p) => (
        <Typography>{p.id}</Typography>
      ))}
    </Box>
  );
}
