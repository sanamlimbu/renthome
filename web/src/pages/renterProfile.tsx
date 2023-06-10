import ArrowBackIosIcon from "@mui/icons-material/ArrowBackIos";
import { Box, Link, Typography, styled } from "@mui/material";
import { blue } from "@mui/material/colors";
import { useNavigate } from "react-router-dom";
import { ReactComponent as CheckedIcon } from "../assets/icons/checked.svg";

export default function RenterProfile() {
  const navigate = useNavigate();
  const link = "rent/renter-profile";
  return (
    <Box>
      <Box
        sx={{
          display: "flex",
          cursor: "pointer",
          color: blue[500],
        }}
        onClick={() => navigate("/me")}
      >
        <ArrowBackIosIcon fontSize="small" sx={{ fontWeight: "bold" }} />
        <Link
          sx={{
            fontWeight: "bold",
            textDecoration: "none",
            "&:hover": {
              textDecoration: "underline",
              color: blue[500],
              opacity: "1",
            },
          }}
          href="/me"
        >
          Back to Overview
        </Link>
      </Box>
      <Typography
        variant="h1"
        sx={{ fontWeight: "400", paddingTop: "1em", paddingBottom: "0.5em" }}
      >
        Renter Profile
      </Typography>
      <Typography>
        Create your Renter Profile once and reuse it for all your applications.
      </Typography>
      <Typography variant="h3" sx={{ paddingTop: "1em", paddingBottom: "1em" }}>
        Personal
      </Typography>
      <Typography sx={{ maxWidth: "36em", marginBottom: "1.5em" }}>
        Details to help property managers validate who you are and assess your
        identity, employment and income.
      </Typography>
      <Box sx={{ display: "flex", flexDirection: "column", gap: "1em" }}>
        <PersonalCard text={"Personal details"} link={"personal-details"} />
        <PersonalCard text={"About me"} link={`${link}/about-me`} />
        <PersonalCard
          text={"Address history"}
          link={`${link}/address-history`}
        />
        <PersonalCard text={"Employment"} link={`${link}/employment`} />
        <PersonalCard text={"Income"} link={`${link}/income`} />
        <PersonalCard
          text={"Identity documents"}
          link={`${link}/identity-documents`}
        />
        <PersonalCard
          text={"Emergency contact"}
          link={`${link}/emergency-contact`}
        />
        <PersonalCard
          text={"Tenant check (recommended"}
          link={`${link}/tenant-check`}
        />
      </Box>
    </Box>
  );
}

const StyledCheckedIcon = styled(CheckedIcon)({
  width: "20px",
  height: "20px",
  verticalAlign: "middle",
});

interface PersonalCardProps {
  text: String;
  link: String;
}

const PersonalCard = ({ text, link }: PersonalCardProps) => {
  const navigate = useNavigate();
  console.log(link);
  return (
    <Typography
      sx={{
        padding: "1em",
        background: "white",
        borderRadius: "12px",
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
        maxWidth: "36em",
        cursor: "pointer",
      }}
      onClick={() => navigate(`${link}`)}
    >
      <span>{text}</span>
      <StyledCheckedIcon />
    </Typography>
  );
};
