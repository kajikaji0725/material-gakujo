import React, { useContext } from "react";
import Box from "@mui/material/Box";
import { Button, Grid, TextField, Typography } from "@mui/material";
import { ApiClientContext } from "../App";
import { Controller, SubmitHandler, useForm } from "react-hook-form";
import { Link } from "react-router-dom";

interface FormInput {
  username: string;
  password: string;
}
export function SignInPage(): JSX.Element {
  const { control, handleSubmit } = useForm<FormInput>();
  const client = useContext(ApiClientContext);
  const onSubmit: SubmitHandler<FormInput> = (data) => {
    client.login(data.username, data.password);
  };

  return (
    <Grid
      container
      alignItems="center"
      justifySelf="center"
      justifyContent="center"
    >
      <Box boxShadow={1} maxWidth={"1200px"}>
        <Typography fontSize="1.2rem">ログイン</Typography>

        <Box
          component="form"
          onSubmit={handleSubmit(onSubmit)}
          sx={{
            width: "100%", // Fix IE 11 issue.
          }}
        >
          <Grid
            container
            alignContent="center"
            alignItems="center"
            justifyContent="center"
            display="flex"
            flexDirection="column"
          >
            <Grid item xs={10}>
              <Controller
                name="username"
                control={control}
                defaultValue=""
                render={({ field }) => (
                  <TextField
                    {...field}
                    required
                    label="学情の username(id)"
                    type="username"
                    margin="normal"
                    style={{ width: "15em" }}
                    fullWidth
                  />
                )}
              />
            </Grid>

            <Grid item>
              <Controller
                name="password"
                control={control}
                defaultValue=""
                render={({ field }) => (
                  <TextField
                    {...field}
                    required
                    label="学情の password"
                    type="password"
                    margin="normal"
                  />
                )}
              />
            </Grid>

            <Grid item textAlign="left">
              <Link
                to="/auth/register"
                style={{
                  textDecoration: "none",
                  color: "#1976d2",
                }}
              >
                <p>新規登録をする</p>
              </Link>
            </Grid>

            <Grid item>
              <Button type="submit" variant="contained" color="primary">
                login
              </Button>
            </Grid>
          </Grid>
        </Box>
      </Box>
    </Grid>
  );
}

/*
<Box
        className={classes.paper}
        onSubmit={handleSubmit(onSubmit)}
        boxShadow={1}
      >
        <Typography variant="h4" align="left">
          ログイン
        </Typography>
        <Box
          component="form"
          m={2}
          sx={{
            width: "100%", // Fix IE 11 issue.
          }}
        >
          <Controller
            name="username"
            control={control}
            defaultValue=""
            render={({ field }) => (
              <TextField
                {...field}
                required
                label="学情の username(id)"
                type="username"
                fullWidth
                margin="normal"
              />
            )}
          />

          <Controller
            name="password"
            control={control}
            defaultValue=""
            render={({ field }) => (
              <TextField
                {...field}
                required
                label="学情の password"
                type="password"
                margin="normal"
                sx={{
                  textAlign: "center",
                }}
              />
            )}
          />

          <Link
            to="/auth/register"
            style={{
              textAlign: "left",
              textDecoration: "none",
              color: "#1976d2",
            }}
          >
            <p>新規登録をする</p>
          </Link>

          <Button type="submit" variant="contained" color="primary">
            login
          </Button>
        </Box>
      </Box>
      */
