package com.airplatform.incidents.service;

import com.airplatform.incidents.domain.Incident;
import com.airplatform.incidents.repository.IncidentRepository;
import com.airplatform.incidents.web.CreateIncidentRequest;
import com.airplatform.incidents.web.IncidentResponse;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
public class IncidentServiceTest {
    @Mock
    IncidentRepository repository;

    @InjectMocks
    IncidentService service;

    @Test
    void create_savesIncidentAndReturnsNewStatus() {
        Incident saved = incidentWith("payments-api", "critical", "NEW", "db timeout");
        when(repository.save(any(Incident.class))).thenReturn(saved);

        IncidentResponse response = service.create(new
                CreateIncidentRequest("payments-api", "critical", "db timeout"));

        assertThat(response.status()).isEqualTo("NEW");
        assertThat(response.service()).isEqualTo("payments-api");
        assertThat(response.severity()).isEqualTo("critical");
        assertThat(response.message()).isEqualTo("db timeout");
        verify(repository).save(any(Incident.class));
    }

    @Test
    void findAll_returnsMappedList() {
        Incident i = incidentWith("auth-service", "warning", "NEW", "slow query");
        when(repository.findAll()).thenReturn(List.of(i));

        List<IncidentResponse> result = service.findAll();

        assertThat(result).hasSize(1);
        assertThat(result.get(0).service()).isEqualTo("auth-service");
    }

    @Test
    void findById_returnsEmptyWhenNotFound() {
        UUID id = UUID.randomUUID();
        when(repository.findById(id)).thenReturn(Optional.empty());

        Optional<IncidentResponse> result = service.findById(id);

        assertThat(result).isEmpty();
    }

    @Test
    void findById_returnsMappedResponseWhenFound() {
        UUID id = UUID.randomUUID();
        Incident i = incidentWith("payments-api", "critical", "NEW", "db timeout");
        when(repository.findById(id)).thenReturn(Optional.of(i));

        Optional<IncidentResponse> result = service.findById(id);

        assertThat(result).isPresent();
        assertThat(result.get().service()).isEqualTo("payments-api");
        assertThat(result.get().status()).isEqualTo("NEW");
    }

    private Incident incidentWith(String service, String severity, String status, String message) {
        Incident i = new Incident();
        i.setServiceName(service);
        i.setSeverity(severity);
        i.setStatus(status);
        i.setMessage(message);
        return i;
    }
}
