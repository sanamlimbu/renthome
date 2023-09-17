import { Close, Search } from "@mui/icons-material";
import {
  Box,
  Button,
  Card,
  Checkbox,
  Dialog,
  Divider,
  FormControl,
  FormControlLabel,
  FormGroup,
  IconButton,
  MenuItem,
  Select,
  TextField,
  Typography,
} from "@mui/material";
import { SetStateAction, useRef, useState } from "react";
import { Controller, SubmitHandler, useForm } from "react-hook-form";
import SelelectedLocationCard from "../components/selectedLocationCard";
import SuggestedLocationCard from "../components/suggestLocationCard";
import { API_ADDRESS } from "../config";

interface IFormInput {
  suburb: string;
  unit: boolean;
  apartment: boolean;
  house: boolean;
  townhouse: boolean;
  all: boolean;
  villa: boolean;
  priceMin: string;
  priceMax: string;
  bedMin: string;
  bedMax: string;
  bathMin: string;
  carMin: string;
}

interface SuggestedLocationsResponse {
  locations: string[];
  total: number;
}

export default function SearchBox() {
  const [filterType, setFilterType] = useState("Buy");
  const [searchType, setSearchType] = useState("Buy");
  const [searchTerm, setSearchTerm] = useState("");
  const [suggestedLocations, setSuggestedLocations] = useState<string[]>([]);
  const [selectedLocations, setSelectedLocations] = useState<string[]>([]);
  const textFieldRef = useRef<HTMLInputElement>();

  const [openFilter, setOpenFilter] = useState(false);
  const { control, handleSubmit, register, reset } = useForm({
    defaultValues: {
      suburb: "",
      unit: false,
      apartment: false,
      villa: false,
      all: true,
      townhouse: false,
      house: false,
      priceMin: "Any",
      priceMax: "Any",
      bedMin: "Any",
      bedMax: "Any",
      bathMin: "Any",
      carMin: "Any",
    },
  });

  const onSubmit: SubmitHandler<IFormInput> = () => {
    setOpenFilter(false);
  };

  const handleSearchTermChange = (event: {
    target: { value: SetStateAction<string> };
  }) => {
    setSearchTerm(event.target.value);
    fetchSuggestedLocations(event.target.value.toString());
  };

  const fetchSuggestedLocations = async (searchTerm: string) => {
    try {
      if (searchTerm === "") {
        setSuggestedLocations([]);
        return;
      }
      const res = await fetch(`${API_ADDRESS}/api/locations`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          search_term: searchTerm,
          ignore_terms: selectedLocations,
        }),
      });
      if (res.ok) {
        const data: SuggestedLocationsResponse = await res.json();
        console.log(data);
        setSuggestedLocations(data.locations);
      }
    } catch (err) {}
  };

  const onSuggestedLocationSelection = (location: string) => {
    if (!selectedLocations.includes(location)) {
      setSelectedLocations([...selectedLocations, location]);
    }
    setSuggestedLocations([]);
    if (textFieldRef.current) {
      textFieldRef.current.value = "";
    }
  };

  const onSuggestedLocationsRemoval = (location: string) => {
    setSelectedLocations(
      selectedLocations.filter((_location) => _location !== location)
    );
  };

  const rentPriceValues = [
    50, 75, 100, 125, 150, 175, 200, 225, 250, 275, 300, 325, 350, 400, 425,
    450, 475, 500,
  ];
  const rentBedValues = [1, 2, 3, 4, 5];
  const rentBathValues = [1, 2, 3, 4, 5];
  const rentCarValues = [1, 2, 3, 4, 5];

  console.log(searchTerm);

  return (
    <Box>
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        <Card sx={{ width: "40rem", padding: "10px", margin: "0px" }}>
          <div
            style={{
              display: "flex",
              justifyContent: "space-around",
              padding: "10px",
            }}
          >
            <Typography
              sx={{
                fontWeight: "600",
                cursor: "pointer",
                "&:hover": {
                  textDecoration: "underline",
                },
                borderBottom: searchType === "Buy" ? "2px solid blue" : "",
                textDecoration: searchType === "Buy" ? "none !important" : "",
              }}
              onClick={() => setSearchType("Buy")}
            >
              Buy
            </Typography>
            <Typography
              sx={{
                fontWeight: "600",
                cursor: "pointer",
                "&:hover": {
                  textDecoration: "underline",
                },
                borderBottom: searchType === "Rent" ? "2px solid blue" : "",
                textDecoration: searchType === "Rent" ? "none !important" : "",
              }}
              onClick={() => setSearchType("Rent")}
            >
              Rent
            </Typography>
            <Typography
              sx={{
                fontWeight: "600",
                cursor: "pointer",
                "&:hover": {
                  textDecoration: "underline",
                },
                borderBottom: searchType === "Sold" ? "2px solid blue" : "",
                textDecoration: searchType === "Sold" ? "none !important" : "",
              }}
              onClick={() => setSearchType("Sold")}
            >
              Sold
            </Typography>
            <Typography
              sx={{
                fontWeight: "600",
                cursor: "pointer",
                "&:hover": {
                  textDecoration: "underline",
                },
                borderBottom: searchType === "Address" ? "2px solid blue" : "",
                textDecoration:
                  searchType === "Address" ? "none !important" : "",
              }}
              onClick={() => setSearchType("Address")}
            >
              Address
            </Typography>
            <Typography
              sx={{
                fontWeight: "600",
                cursor: "pointer",
                "&:hover": {
                  textDecoration: "underline",
                },
                borderBottom: searchType === "Agents" ? "2px solid blue" : "",
                textDecoration:
                  searchType === "Agents" ? "none !important" : "",
              }}
              onClick={() => setSearchType("Agents")}
            >
              Agents
            </Typography>
          </div>
          <Divider />
          <Box
            sx={{
              display: "flex",
              paddingTop: "0.5em",
              paddingBottom: "0.5em",
              gap: "12px",
              alignItems: "center",
            }}
          >
            <Box
              width="100%"
              sx={{
                display: "flex",
                alignItems: "center",
                "&:hover": {
                  background:
                    selectedLocations.length > 0 ? "" : "rgb(246,245,247)",
                },
                padding: "0.5em",
                borderRadius: "6px",
                gap: "2px",
                flexWrap: selectedLocations.length > 0 ? "wrap" : "",
              }}
            >
              <Search color="action" />
              {selectedLocations.map((location, index) => (
                <SelelectedLocationCard
                  key={index}
                  location={location}
                  handleCancel={onSuggestedLocationsRemoval}
                />
              ))}
              <TextField
                fullWidth
                size="small"
                placeholder={
                  searchType === "Address"
                    ? "Search by address"
                    : "Search suburb, postcode or state"
                }
                {...register("suburb")}
                onChange={handleSearchTermChange}
                sx={{
                  "& fieldset": { border: "none" },
                }}
                inputRef={textFieldRef}
              />
            </Box>

            {searchType !== "Address" && searchType !== "Agents" && (
              <Button
                sx={{
                  textTransform: "none",
                }}
                variant="outlined"
                onClick={() => setOpenFilter(true)}
              >
                Filters
              </Button>
            )}
            <Button
              variant="contained"
              color="info"
              sx={{ textTransform: "none" }}
              type="submit"
              onClick={handleSubmit(onSubmit)}
            >
              Search
            </Button>
          </Box>
          <div>
            {suggestedLocations.length > 0 && (
              <div>
                <Divider />
                <Typography
                  sx={{
                    marginLeft: "0.5em",
                    marginTop: "0.5em",
                    marginBottom: "0.5em",
                    fontWeight: "600",
                  }}
                >
                  Suggested locations
                </Typography>
                {suggestedLocations.map((location, index) => (
                  <SuggestedLocationCard
                    key={index}
                    location={location}
                    handleSelect={onSuggestedLocationSelection}
                  />
                ))}
              </div>
            )}
          </div>
        </Card>
      </Box>

      {/**
       * Filter dialog box
       */}
      <Dialog open={openFilter} onClose={() => setOpenFilter(false)}>
        <form onSubmit={handleSubmit(onSubmit)}>
          <Typography
            sx={{
              position: "sticky",
              top: "0",
              fontWeight: "600",
              display: "flex",
              justifyContent: "space-between",
              alignItems: "center",
              padding: "20px 40px 20px 40px",
              background: "white",
              zIndex: "999",
              borderBottom: "1px solid lightgray",
            }}
          >
            <span style={{ fontWeight: "bold", fontSize: "20px" }}>
              Filters
            </span>
            <IconButton onClick={() => setOpenFilter(false)}>
              <Close />
            </IconButton>
          </Typography>
          <div
            style={{
              display: "flex",
              justifyContent: "space-around",
              padding: "20px",
            }}
          >
            <Typography
              sx={{
                cursor: "pointer",
                fontWeight: "600",
                "&:hover": {
                  textDecoration: "underline",
                },
                borderBottom: filterType === "Buy" ? "2px solid blue" : "",
                textDecoration: filterType === "Buy" ? "none !important" : "",
              }}
              onClick={() => setFilterType("Buy")}
            >
              Buy
            </Typography>
            <Typography
              sx={{
                cursor: "pointer",
                fontWeight: "600",
                "&:hover": {
                  textDecoration: "underline",
                },
                borderBottom: filterType === "Rent" ? "2px solid blue" : "",
                textDecoration: filterType === "Rent" ? "none !important" : "",
              }}
              onClick={() => setFilterType("Rent")}
            >
              Rent
            </Typography>
            <Typography
              sx={{
                cursor: "pointer",
                fontWeight: "600",
                "&:hover": {
                  textDecoration: "underline",
                },
                borderBottom: filterType === "Sold" ? "2px solid blue" : "",
                textDecoration: filterType === "Sold" ? "none !important" : "",
              }}
              onClick={() => setFilterType("Sold")}
            >
              Sold
            </Typography>
          </div>
          <Divider />
          <div
            style={{
              display: "flex",
              flexDirection: "column",
              padding: "40px",
              gap: "20px",
            }}
          >
            <Typography sx={{ fontWeight: "600" }}>Property type</Typography>
            <FormGroup
              sx={{
                display: "flex",
                flexDirection: "row",
                gap: "10px",
              }}
            >
              <FormControlLabel
                control={<Checkbox defaultChecked />}
                label="All types"
                {...register("all")}
              />
              <FormControlLabel
                {...register("apartment")}
                control={<Checkbox />}
                label="Apartment"
              />
              <FormControlLabel
                control={<Checkbox />}
                label="Unit"
                {...register("unit")}
              />
              <FormControlLabel
                control={<Checkbox />}
                label="Villa"
                {...register("villa")}
              />
              <FormControlLabel
                control={<Checkbox />}
                label="House"
                {...register("house")}
              />
              <FormControlLabel
                control={<Checkbox />}
                label="Townhouse"
                {...register("townhouse")}
              />
            </FormGroup>
            <Divider />
            <Typography sx={{ fontWeight: "600" }}>Price</Typography>
            <div
              style={{
                display: "flex",
                gap: "20px",
                justifyContent: "space-between",
              }}
            >
              <FormControl fullWidth {...register("priceMin")}>
                <Controller
                  name="priceMin"
                  control={control}
                  render={({ field }) => (
                    <Select {...field}>
                      <MenuItem value={"Any"}>Any</MenuItem>
                      {rentPriceValues.map((value) => (
                        <MenuItem key={"rent-price-min" + value} value={value}>
                          {"$" + value}
                        </MenuItem>
                      ))}
                    </Select>
                  )}
                />
              </FormControl>
              <FormControl fullWidth>
                <Controller
                  name="priceMax"
                  control={control}
                  render={({ field }) => (
                    <Select {...field}>
                      <MenuItem value={"Any"}>Any</MenuItem>
                      {rentPriceValues.map((value) => (
                        <MenuItem key={"rent-price-max" + value} value={value}>
                          {"$" + value}
                        </MenuItem>
                      ))}
                    </Select>
                  )}
                />
              </FormControl>
            </div>
            <Divider />
            <Typography sx={{ fontWeight: "600" }}>Bedrooms</Typography>
            <div
              style={{
                display: "flex",
                gap: "20px",
                justifyContent: "space-between",
              }}
            >
              <FormControl fullWidth>
                <Controller
                  name="bedMin"
                  control={control}
                  render={({ field }) => (
                    <Select {...field}>
                      <MenuItem value={"Any"}>Any</MenuItem>
                      {rentBedValues.map((value) => (
                        <MenuItem key={"rent-bed-min" + value} value={value}>
                          {value}
                        </MenuItem>
                      ))}
                    </Select>
                  )}
                />
              </FormControl>
              <FormControl fullWidth>
                <Controller
                  name="bedMax"
                  control={control}
                  render={({ field }) => (
                    <Select {...field}>
                      <MenuItem value={"Any"}>Any</MenuItem>
                      {rentBedValues.map((value) => (
                        <MenuItem key={"rent-bed-max" + value} value={value}>
                          {value}
                        </MenuItem>
                      ))}
                    </Select>
                  )}
                />
              </FormControl>
            </div>
            <Divider />
            <Typography sx={{ fontWeight: "600" }}>Bathrooms</Typography>
            <div
              style={{
                display: "flex",
                gap: "20px",
                justifyContent: "space-between",
              }}
            >
              <FormControl fullWidth>
                <Controller
                  name="bathMin"
                  control={control}
                  render={({ field }) => (
                    <Select {...field}>
                      <MenuItem value={"Any"}>Any</MenuItem>
                      {rentBathValues.map((value) => (
                        <MenuItem key={"rent-bath-min" + value} value={value}>
                          {value + "+"}
                        </MenuItem>
                      ))}
                    </Select>
                  )}
                />
              </FormControl>
              <div style={{ width: "100%" }}></div>
            </div>
            <Divider />
            <Typography sx={{ fontWeight: "600" }}>Car spaces</Typography>
            <div
              style={{
                display: "flex",
                gap: "20px",
                justifyContent: "space-between",
              }}
            >
              <FormControl fullWidth>
                <Controller
                  name="carMin"
                  control={control}
                  render={({ field }) => (
                    <Select {...field}>
                      <MenuItem value={"Any"}>Any</MenuItem>
                      {rentCarValues.map((value) => (
                        <MenuItem key={"rent-car-min" + value} value={value}>
                          {value + "+"}
                        </MenuItem>
                      ))}
                    </Select>
                  )}
                />
              </FormControl>

              <div style={{ width: "100%" }}></div>
            </div>
          </div>
          <Typography
            sx={{
              position: "sticky",
              bottom: "0",
              background: "white",
              display: "flex",
              justifyContent: "space-between",
              padding: "20px 40px 20px 40px",
              borderTop: "1px solid lightgray",
            }}
          >
            <Button
              sx={{ textTransform: "none", fontWeight: "bold" }}
              onClick={() => {
                reset(
                  (formValues) => ({
                    ...formValues,
                  }),
                  {
                    keepDefaultValues: true,
                  }
                );
              }}
            >
              Clear filters
            </Button>
            <Button
              sx={{ textTransform: "none", fontWeight: "bold" }}
              variant="contained"
              type="submit"
            >
              Search
            </Button>
          </Typography>
        </form>
      </Dialog>
    </Box>
  );
}
