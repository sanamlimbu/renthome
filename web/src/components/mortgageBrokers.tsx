import { Box, Typography } from "@mui/material";
import MortgageBrokersCard, {
  MortgageBrokersCardProps,
} from "./mortgageBrokersCard";

import MortgageChoice from "../assets/mortgage.svg";

export default function MortgageBrokers() {
  return (
    <Box>
      <Typography variant="h3" marginBottom={"1em"}>
        Local mortgage brokers
      </Typography>
      <div style={{ display: "flex", gap: "1em", flexWrap: "wrap" }}>
        {mortgageBrokers.map((card, i) => (
          <MortgageBrokersCard
            key={i}
            imgSrc={card.imgSrc}
            companyLogoSrc={card.companyLogoSrc}
            companyName={card.companyName}
            name={card.name}
            locationCity={card.locationCity}
          />
        ))}
      </div>
    </Box>
  );
}

const mortgageBrokers: MortgageBrokersCardProps[] = [
  {
    companyLogoSrc: MortgageChoice,
    companyName: "",
    name: "Peter Keenan",
    locationCity: "Perth",
    imgSrc: "",
  },
  {
    companyLogoSrc: MortgageChoice,
    companyName: "",
    name: "Bevan O'Farrell",
    locationCity: "Perth",
    imgSrc: "",
  },
  {
    companyLogoSrc: MortgageChoice,
    companyName: "",
    name: "John W Vodanovic",
    locationCity: "North Fremantle",
    imgSrc: "",
  },
  {
    companyLogoSrc: MortgageChoice,
    companyName: "",
    name: "Chris Brown",
    locationCity: "Mount Lawley",
    imgSrc: "",
  },
];
