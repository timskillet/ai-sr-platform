package com.airplatform.incidents.web;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;

import java.util.UUID;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;
@SpringBootTest
@AutoConfigureMockMvc
public class IncidentControllerTest {
    @Autowired
    MockMvc mockMvc;

    @Autowired
    ObjectMapper objectMapper;

    @Test
    void createIncident_returns201WithCorrectFields() throws Exception {
        String body = """
                {"service":"payments-api","severity":"critical",
                "message":"db timeout"}
                """;

        mockMvc.perform(post("/incidents")
                .contentType(MediaType.APPLICATION_JSON)
                .content(body))
            .andExpect(status().isCreated())
            .andExpect(jsonPath("$.id").exists())
            .andExpect(jsonPath("$.service").value("payments-api"))
            .andExpect(jsonPath("$.severity").value("critical"))
            .andExpect(jsonPath("$.status").value("NEW"))
            .andExpect(jsonPath("$.message").value("db timeout"))
            .andExpect(jsonPath("$.createdAt").exists());
    }

    @Test
    void getIncidents_includesCreatedIncident() throws Exception {
        String body = """
                {"service":"auth-service","severity":"warning",
                "message":"slow query"}
                """;

        mockMvc.perform(post("/incidents")
                .contentType(MediaType.APPLICATION_JSON)
                .content(body))
            .andExpect(status().isCreated());

        mockMvc.perform(get("/incidents"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$").isArray())
                .andExpect(jsonPath("$[?(@.service == 'auth-service')]").exists());
    }

    @Test
    void getIncidentById_returns404ForUnknownId() throws Exception {
        mockMvc.perform(get("/incidents/" + UUID.randomUUID()))
                .andExpect(status().isNotFound());
    }

    @Test
    void getIncidentById_returnsIncidentWhenFound() throws Exception {
        String body = """
                {"service":"payment-service","severity":"critical",
                "message":"timeout"}
                """;

        String responseBody = mockMvc.perform(post("/incidents")
                .contentType(MediaType.APPLICATION_JSON)
                .content(body))
                .andExpect(status().isCreated())
                .andReturn().getResponse().getContentAsString();

        String id = objectMapper.readTree(responseBody).get("id").asText();

        mockMvc.perform(get("/incidents/" + id))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.id").value(id))
                .andExpect(jsonPath("$.service").value("payment-service"))
                .andExpect(jsonPath("$.status").value("NEW"));
    }
}
