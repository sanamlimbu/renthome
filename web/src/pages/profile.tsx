import { Alert, Box, Button, Snackbar, Typography } from "@mui/material";
import { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";
import { ReactComponent as AccountSettings } from "../assets/profile/account-settings.svg";
import { ReactComponent as Finances } from "../assets/profile/finances.svg";
import { ReactComponent as MyProfile } from "../assets/profile/my-profile.svg";
import { ReactComponent as MySavedProperties } from "../assets/profile/my-saved-properties.svg";
import { ReactComponent as PrivateLandlord } from "../assets/profile/private-landlord.svg";
import { ReactComponent as RentalApplications } from "../assets/profile/rental-applications.svg";
import { ReactComponent as RenterProfile } from "../assets/profile/renter-profile.svg";
import { ReactComponent as SavedSearches } from "../assets/profile/saved-searches.svg";
import { ReactComponent as TrackProperty } from "../assets/profile/track-property.svg";
import { API_ADDRESS } from "../config";
import { UserContext } from "../context/user";
import {
  getTokenFromLocalStorage,
  removeTokenFromLocalStorage,
} from "../helpers/auth";

export default function ProfilePage() {
  const { user, setUser } = useContext(UserContext);
  const [error, setError] = useState("");
  const [openSnackbar, setOpenSnackbar] = useState(false);

  const handleSnackbarClose = () => {
    setOpenSnackbar(false);
  };

  const handleLogout = async () => {
    try {
      const res = await fetch(`${API_ADDRESS}/auth/logout`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${getTokenFromLocalStorage()}`,
        },
        body: JSON.stringify({ user_id: `${user?.id}` }),
      });
      if (!res.ok) {
        const data = await res.json();
        setError(data.message);
        setOpenSnackbar(true);
      } else {
        removeTokenFromLocalStorage();
        setUser(undefined);
        setOpenSnackbar(false);
      }
    } catch (error) {
      setOpenSnackbar(true);
    }
  };
  return (
    <Box
      sx={{
        width: "100%",
        maxWidth: "60rem",
      }}
    >
      <Box
        sx={{
          display: "flex",
          justifyContent: "space-between",
          marginBottom: "2em",
        }}
      >
        <Typography variant="h1"> Profile</Typography>
        <Button
          variant="contained"
          sx={{ fontWeight: "bold" }}
          onClick={handleLogout}
        >
          Log out
        </Button>
      </Box>
      <Box
        sx={{
          display: "flex",
          flexWrap: "wrap",
          justifyContent: "space-between",
          gap: "1em",
        }}
      >
        <ProfileCard
          icon={TrackProperty}
          title={"Track your property"}
          shortDescription={
            "Stay up to date with your home or properties you own"
          }
          link={"property/my-property"}
        />

        <ProfileCard
          icon={MySavedProperties}
          title={"My saved properties"}
          shortDescription={
            "View open times and auctions for properties you've saved."
          }
          link={"collections/saved-properties/"}
        />
        <ProfileCard
          icon={SavedSearches}
          title={"Saved searches & alerts"}
          shortDescription={
            "View your saved searches and configure their alerts"
          }
          link={"saved-searches/"}
        />
        <ProfileCard
          icon={RenterProfile}
          title={"Renter Profile"}
          shortDescription={"Create or update your Renter Profile"}
          link={"rent/renter-profile/"}
        />
        <ProfileCard
          icon={RentalApplications}
          title={"Rental applications"}
          shortDescription={
            "Track the status and view your rental applications"
          }
          link={"rent/applications"}
        />
        <ProfileCard
          icon={Finances}
          title={"My finances"}
          shortDescription={
            "Financial tools to help you make better property decisions."
          }
          link={"my-finances/"}
        />
        <ProfileCard
          icon={AccountSettings}
          title={"Account settings"}
          shortDescription={
            "Manage your password, email subscriptions and privacy settings."
          }
          link={"my-real-estate/account/"}
        />
        <ProfileCard
          icon={MyProfile}
          title={"My profile"}
          shortDescription={"Manage your personal details and property needs."}
          link={"profile/"}
        />
        <ProfileCard
          icon={PrivateLandlord}
          title={"My rental listing"}
          shortDescription={"Create and manage your rental property listings."}
          link={"advertise-property-for-rent/manage-listings"}
        />
      </Box>
      <Snackbar
        open={openSnackbar}
        autoHideDuration={6000}
        onClose={handleSnackbarClose}
        anchorOrigin={{
          vertical: "bottom",
          horizontal: "center",
        }}
      >
        <Alert onClose={handleSnackbarClose} severity="error">
          {error || "Something went wrong, please try again."}
        </Alert>
      </Snackbar>
    </Box>
  );
}

interface ProfileCardProps {
  icon: React.ComponentType<React.SVGProps<SVGSVGElement>>;
  title: String;
  shortDescription: String;
  link: String;
}
function ProfileCard({
  icon: Icon,
  title,
  shortDescription,
  link,
}: ProfileCardProps) {
  const navigate = useNavigate();
  return (
    <Box
      sx={{
        padding: "1em",
        background: "white",
        borderRadius: "12px",
        boxShadow: "0px 3px 6px rgba(0, 0, 0, 0.16)",
        cursor: "pointer",
      }}
      onClick={() => navigate(`/${link}`)}
    >
      <Icon />
      <Typography
        variant="h4"
        style={{ paddingTop: "0.5em", paddingBottom: "0.5em" }}
      >
        {title}
      </Typography>
      <Typography sx={{ maxWidth: "16em" }}>{shortDescription}</Typography>
    </Box>
  );
}
