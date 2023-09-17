import { Box, Typography } from "@mui/material";
import Banner from "../assets/banner.avif";
import Explore from "../components/explore";
import MortgageBrokers from "../components/mortgageBrokers";
import News from "../components/news";
import SearchBox from "../components/searchBox";
import { FilterType, SearchType } from "../types/enums";

export default function SoldPage() {
  return (
    <Box>
      <Box
        sx={{
          position: "relative",
          zIndex: 2,
        }}
      >
        <Box
          sx={{
            position: "absolute",
            left: 0,
            right: 0,
            margin: "0 auto",
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            justifyContent: "center",
          }}
        >
          <Typography
            variant="h3"
            sx={{
              marginTop: "2rem",
              marginBottom: "1.2rem",
            }}
          >
            Search sold properties
          </Typography>
          <SearchBox
            filterType={FilterType.Sold}
            searchType={SearchType.Sold}
          />
        </Box>
        <Box
          sx={{
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            maxHeight: "30rem",
            marginLeft: "2rem",
            marginRight: "2rem",
            marginTop: "2rem",
            borderRadius: "20px",
            overflow: "hidden",
          }}
        >
          <img
            src={Banner}
            alt="Perth suburb"
            style={{ maxWidth: "100%", maxHeight: "100%" }}
          />
        </Box>
      </Box>

      <Box
        sx={{
          marginLeft: "10rem",
          marginRight: "10rem",
          marginTop: "2rem",
          zIndex: 1,
        }}
      >
        <Box sx={{ marginBottom: "3em" }}>
          <Explore />
        </Box>
        <Box sx={{ marginBottom: "3em" }}>
          <News />
        </Box>
        <Box sx={{ marginBottom: "3em" }}>
          <MortgageBrokers />
        </Box>
      </Box>
    </Box>
  );
}
