import { Box, Button, Typography } from "@mui/material";
import { useState } from "react";
import Auctions from "../assets/auctions.avif";
import CalculatorsBorrow from "../assets/calculators-borrow.avif";
import Calculators from "../assets/calculators.avif";
import ExploreProperty from "../assets/explore.avif";
import Flatmates from "../assets/flatmates.avif";
import GuidesBuying from "../assets/guides-buying.avif";
import GuidesSelling from "../assets/guides-selling.avif";
import HomeLoan from "../assets/homeloan.avif";
import Market from "../assets/market.avif";
import Rent from "../assets/rent.avif";
import Tenant from "../assets/tenant.avif";
import TrackProperty from "../assets/track-property.avif";
import { ExploreType } from "../types/enums";
import ExploreCard, { ExploreCardProps } from "./exploreCard";

export default function Explore() {
  const [exploreType, setExploreType] = useState<ExploreType>(
    ExploreType.Buying
  );
  return (
    <Box sx={{ display: "flex", flexDirection: "column" }}>
      <Typography variant="h3">Explore all things property</Typography>
      <div style={{ marginTop: "1em", marginBottom: "1em" }}>
        <Button
          variant="outlined"
          sx={{
            borderRadius: "22px",
            fontWeight: "bold",
            marginRight: "1em",
            border:
              exploreType === ExploreType.Buying ? "2px solid !important" : "",
          }}
          onClick={() => setExploreType(ExploreType.Buying)}
        >
          Buying
        </Button>
        <Button
          variant="outlined"
          sx={{
            borderRadius: "22px",
            fontWeight: "bold",
            marginRight: "1em",
            border:
              exploreType === ExploreType.Renting ? "2px solid !important" : "",
          }}
          onClick={() => setExploreType(ExploreType.Renting)}
        >
          Renting
        </Button>
        <Button
          variant="outlined"
          sx={{
            borderRadius: "22px",
            fontWeight: "bold",
            marginRight: "1em",
            border:
              exploreType === ExploreType.Selling ? "2px solid !important" : "",
          }}
          onClick={() => setExploreType(ExploreType.Selling)}
        >
          Selling
        </Button>
        <Button
          variant="outlined"
          sx={{
            borderRadius: "22px",
            fontWeight: "bold",
            marginRight: "1em",
            border:
              exploreType === ExploreType.Researching
                ? "2px solid !important"
                : "",
          }}
          onClick={() => setExploreType(ExploreType.Researching)}
        >
          Researching
        </Button>
      </div>
      <div style={{ display: "flex", gap: "1em", flexWrap: "wrap" }}>
        {exploreType === ExploreType.Buying &&
          buyingCards.map((card, i) => (
            <ExploreCard
              key={i}
              title={card.title}
              shortDescription={card.shortDescription}
              linkText={card.linkText}
              imgSrc={card.imgSrc}
            />
          ))}
        {exploreType === ExploreType.Renting &&
          rentingCards.map((card, i) => (
            <ExploreCard
              key={i}
              title={card.title}
              shortDescription={card.shortDescription}
              linkText={card.linkText}
              imgSrc={card.imgSrc}
            />
          ))}
        {exploreType === ExploreType.Selling &&
          sellingCards.map((card, i) => (
            <ExploreCard
              key={i}
              title={card.title}
              shortDescription={card.shortDescription}
              linkText={card.linkText}
              imgSrc={card.imgSrc}
            />
          ))}
        {exploreType === ExploreType.Researching &&
          researchingCards.map((card, i) => (
            <ExploreCard
              key={i}
              title={card.title}
              shortDescription={card.shortDescription}
              linkText={card.linkText}
              imgSrc={card.imgSrc}
            />
          ))}
      </div>
    </Box>
  );
}

const buyingCards: ExploreCardProps[] = [
  {
    title: "Get estimated property prices with a realEstimate™",
    shortDescription:
      "See how much your property's worth whether you own it or want to buy it.",
    imgSrc: TrackProperty,
    linkText: "Check property values",
  },
  {
    title: "Need help with a mortgage?",
    shortDescription: "Compare your finance options to make an informed call.",
    imgSrc: HomeLoan,
    linkText: "Explore home loans",
  },
  {
    title: "Explore suburbs profile",
    shortDescription:
      "Check out different suburb profiles and find one that's right for you.",
    imgSrc: ExploreProperty,
    linkText: "Research suburbs",
  },
];
const rentingCards: ExploreCardProps[] = [
  {
    title: "Set up your renter profile",
    shortDescription:
      "Are you a tenant looking for a new place? Get yourself ready to apply online.",
    imgSrc: Rent,
    linkText: "Create a profile",
  },
  {
    title: "Find a tenant",
    shortDescription:
      "List your property for rent through an agent or on your own.",
    imgSrc: Tenant,
    linkText: "Advertise your rental property",
  },
  {
    title: "Looking for a flatmate?",
    shortDescription:
      "Move into a new sharehouses, or find someone to join yours.",
    imgSrc: Flatmates,
    linkText: "Head to flatmates.com.au",
  },
];

const sellingCards: ExploreCardProps[] = [
  {
    title: "Check this week's sold prices",
    shortDescription:
      "Stay in loop with the latest clearance rates and recent sales.",
    imgSrc: Auctions,
    linkText: "See auctions results",
  },
  {
    title: "Decide whether to sell or stay",
    shortDescription: "Wondering if now's good time to sell? See your options.",
    imgSrc: GuidesSelling,
    linkText: "Explore home loans",
  },
  {
    title: "Think about refinancing?",
    shortDescription: "Calculate your equity weigh up your financial options.",
    imgSrc: Calculators,
    linkText: "Use the refinance calculator",
  },
];

const researchingCards: ExploreCardProps[] = [
  {
    title: "Delve into market data",
    shortDescription:
      "Stay on top of real estate trends written by economists and property experts.",
    imgSrc: Market,
    linkText: "Research the property market",
  },
  {
    title: "Real estate how to's",
    shortDescription: "Read up on the ins and outs of any property process.",
    imgSrc: GuidesBuying,
    linkText: "Explore property guides",
  },
  {
    title: "Get your finances in order",
    shortDescription:
      "Use our calculators to understand your financial position and options.",
    imgSrc: CalculatorsBorrow,
    linkText: "Browse home loan calculators",
  },
];
