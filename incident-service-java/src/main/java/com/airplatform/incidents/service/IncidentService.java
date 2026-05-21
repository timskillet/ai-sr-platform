package com.airplatform.incidents.service;

import com.airplatform.incidents.domain.Incident;
import com.airplatform.incidents.repository.IncidentRepository;
import com.airplatform.incidents.web.CreateIncidentRequest;
import com.airplatform.incidents.web.IncidentResponse;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Service
public class IncidentService {
    private final IncidentRepository repository;

    public IncidentService(IncidentRepository repository) {
        this.repository = repository;
    }

    public IncidentResponse create(CreateIncidentRequest request) {
        Incident incident = new Incident();
        incident.setServiceName(request.service());
        incident.setSeverity(request.severity());
        incident.setMessage(request.message());
        incident.setStatus("NEW");
        return toResponse(repository.save(incident));
    }

    public List<IncidentResponse> findAll() {
        return repository.findAll().stream()
                .map(this::toResponse)
                .toList();
    }

    public Optional<IncidentResponse> findById(UUID id) {
        return repository.findById(id).map(this::toResponse);
    }

    private IncidentResponse toResponse(Incident incident) {
        return new IncidentResponse(
                incident.getId(),
                incident.getServiceName(),
                incident.getSeverity(),
                incident.getStatus(),
                incident.getMessage(),
                incident.getCreatedAt()
        );
    }
}
