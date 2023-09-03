import { Box, Typography } from "@mui/material";
import NewsOne from "../assets/news1.jpg";
import NewsTwo from "../assets/news2.jpg";
import NewsThree from "../assets/news3.jpg";
import NewsFour from "../assets/news4.jpeg";

import NewsCard, { NewsCardProps } from "./newsCard";

export default function News() {
  return (
    <Box>
      <Typography variant="h3" marginBottom={"1em"}>
        Lastest property news
      </Typography>
      <div style={{ display: "flex", gap: "1em", flexWrap: "wrap" }}>
        {news.map((card, i) => (
          <NewsCard
            key={i}
            title={card.title}
            readTime={card.readTime}
            imgSrc={card.imgSrc}
          />
        ))}
      </div>
    </Box>
  );
}

const news: NewsCardProps[] = [
  {
    title:
      "The PropTrack Housing Affordability Index: a better way to measure affordability",
    imgSrc: NewsOne,
    readTime: "3 mins read",
  },
  {
    title: "PropTrack Housing Affordability Report - 2023",
    imgSrc: NewsTwo,
    readTime: "4 mins read",
  },
  {
    title: "Housing affordability reaches three decade low",
    imgSrc: NewsThree,
    readTime: "5 mins read",
  },
  {
    title:
      "Carlton North: Family behind Lygon St eatery Tiamo list their two homes for sale ",
    imgSrc: NewsFour,
    readTime: "2 mins read",
  },
];
