import { DataGrid, GridCellParams, GridColDef } from "@material-ui/data-grid";
import { Container } from "@mui/material";
import { Box } from "@mui/system";
import React, { useContext, useEffect, useState } from "react";
import { Seiseki } from "../api/seiseki";
import { ApiClientContext } from "../App";

const columns: GridColDef[] = [
  {
    field: "SubjectName",
    headerName: "科目名",
    sortable: false,
    disableColumnMenu: true,
    width: 250,
  },
  { field: "TeacherName", headerName: "教員名", sortable: false, width: 100 },
  {
    field: "SubjectDistinction",
    headerName: "必/選必/選",
    sortable: false,
    width: 150,
  },
  { field: "SubjectType", headerName: "科目区分", sortable: false, width: 120 },
  {
    field: "Credit",
    headerName: "単位数",
    width: 100,
    disableColumnMenu: true,
  },
  { field: "Grade", headerName: "評価", sortable: false, width: 90 },
  { field: "Score", headerName: "得点", disableColumnMenu: true, width: 85 },
  { field: "GP", headerName: "GP", disableColumnMenu: true, width: 80 },
  { field: "Year", headerName: "年度", width: 105 },
  {
    field: "Date",
    headerName: "報告日",
    width: 130,
    renderCell: (params: GridCellParams) => {
      const date = new Date(params.value as string);
      return `${date.getFullYear()}-${date.getMonth() + 1}-${date.getUTCDay()}`;
    },
  },
];

export function SeisekiPage(): JSX.Element {
  const apiClient = useContext(ApiClientContext);
  const [seisekis, setSeisekis] = useState<Seiseki[]>(new Array(0));

  useEffect(() => {
    const fetchSeisekis = async (): Promise<Seiseki[]> => {
      return await apiClient.fetchSeisekis();
    };
    fetchSeisekis().then((seisekis) => setSeisekis(seisekis));
    return () => {
      console.log(seisekis);
    };
  }, []);

  return (
    <Container component="main" maxWidth="xl">
      <Box component="div" width="100%" height="800px">
        <DataGrid rows={seisekis} columns={columns} disableColumnSelector />
      </Box>
    </Container>
  );
}
