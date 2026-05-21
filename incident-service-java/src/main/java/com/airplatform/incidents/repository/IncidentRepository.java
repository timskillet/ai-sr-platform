package com.airplatform.incidents.repository;

import com.airplatform.incidents.domain.Incident;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.UUID;

public interface IncidentRepository extends JpaRepository<Incident, UUID> {
}
